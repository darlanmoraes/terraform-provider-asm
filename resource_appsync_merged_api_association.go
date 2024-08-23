package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppSyncMergedApiAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceAppSyncMergedApiAssociationCreate,
		Read:   resourceAppSyncMergedApiAssociationRead,
		Update: resourceAppSyncMergedApiAssociationUpdate,
		Delete: resourceAppSyncMergedApiAssociationDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"merged_api_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_api_association_config": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"merge_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "AUTO_MERGE",
						},
					},
				},
			},
			"source_api_identifier": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAppSyncMergedApiAssociationCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.AssociateMergedGraphqlApiInput{
		Description:                aws.String(d.Get("description").(string)),
		MergedApiIdentifier:        aws.String(d.Get("merged_api_identifier").(string)),
		SourceApiAssociationConfig: expandSourceApiAssociationConfig(d.Get("source_api_association_config").([]interface{})),
		SourceApiIdentifier:        aws.String(d.Get("source_api_identifier").(string)),
	}

	result, err := client.AppSync.AssociateMergedGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId(aws.StringValue(result.SourceApiAssociation.AssociationId))

	return resourceAppSyncMergedApiAssociationRead(d, meta)
}

func resourceAppSyncMergedApiAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.GetSourceApiAssociationInput{
		AssociationId:       aws.String(d.Id()),
		MergedApiIdentifier: aws.String(d.Get("merged_api_identifier").(string)),
	}
	result, err := client.AppSync.GetSourceApiAssociation(input)
	if err != nil {
		return err
	}

	d.Set("description", result.SourceApiAssociation.Description)
	d.Set("merged_api_identifier", result.SourceApiAssociation.MergedApiId)
	d.Set("source_api_association_config", flattenSourceApiAssociationConfig(result.SourceApiAssociation.SourceApiAssociationConfig))
	d.Set("source_api_identifier", result.SourceApiAssociation.SourceApiId)

	return nil
}

func resourceAppSyncMergedApiAssociationUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.UpdateSourceApiAssociationInput{
		AssociationId:       aws.String(d.Id()),
		MergedApiIdentifier: aws.String(d.Get("merged_api_identifier").(string)),
	}

	if d.HasChange("description") {
		input.Description = aws.String(d.Get("description").(string))
	}

	if d.HasChange("source_api_association_config") {
		input.SourceApiAssociationConfig = expandSourceApiAssociationConfig(d.Get("source_api_association_config").([]interface{}))
	}

	_, err := client.AppSync.UpdateSourceApiAssociation(input)
	if err != nil {
		return err
	}

	return resourceAppSyncMergedApiAssociationRead(d, meta)
}

func resourceAppSyncMergedApiAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.DisassociateMergedGraphqlApiInput{
		AssociationId:       aws.String(d.Id()),
		SourceApiIdentifier: aws.String(d.Get("source_api_identifier").(string)),
	}

	_, err := client.AppSync.DisassociateMergedGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

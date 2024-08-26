package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAsmAppSyncStartSchemaMerge() *schema.Resource {
	return &schema.Resource{
		Create: resourceAsmAppSyncStartSchemaMergeCreate,
		Read:   schema.Noop,
		Update: resourceAsmAppSyncStartSchemaMergeUpdate,
		Delete: schema.Noop,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"association_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"merged_api_id": {
				Type:     schema.TypeString,
				Required: true,
			},
			"timeout_seconds": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAsmAppSyncStartSchemaMergeCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	associationId := d.Get("association_id").(string)
	mergedApiId := d.Get("merged_api_id").(string)

	input := &appsync.StartSchemaMergeInput{
		AssociationId:       aws.String(associationId),
		MergedApiIdentifier: aws.String(mergedApiId),
	}

	_, err := client.AppSync.StartSchemaMerge(input)
	if err != nil {
		return err
	}

	status, err := wait("MERGE_SUCCESS", d, meta)
	if err != nil {
		return err
	}

	d.SetId(fmt.Sprintf("%s_%s", associationId, mergedApiId))
	d.Set("status", status)

	return nil
}

func resourceAsmAppSyncStartSchemaMergeUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAsmAppSyncStartSchemaMergeCreate(d, meta)
}

func wait(expectedStatus string, d *schema.ResourceData, meta interface{}) (string, error) {
	associationId := d.Get("association_id").(string)
	mergedApiId := d.Get("merged_api_id").(string)
	timeoutSeconds := d.Get("timeout_seconds").(int)

	client := meta.(*AWSClient)
	for i := 0; i < timeoutSeconds; i++ {
		input := &appsync.GetSourceApiAssociationInput{
			AssociationId:       aws.String(associationId),
			MergedApiIdentifier: aws.String(mergedApiId),
		}

		result, err := client.AppSync.GetSourceApiAssociation(input)
		if err != nil {
			return "", err
		}

		status := aws.StringValue(result.SourceApiAssociation.SourceApiAssociationStatus)
		if status == expectedStatus {
			return status, nil
		}

		time.Sleep(1 * time.Second)
	}

	return "", fmt.Errorf("schema merge did not succeed after %d attempts", timeoutSeconds)
}

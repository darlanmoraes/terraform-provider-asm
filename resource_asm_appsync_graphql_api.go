package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAsmAppSyncGraphQLApi() *schema.Resource {
	return &schema.Resource{
		Create: resourceAsmAppSyncGraphQLApiCreate,
		Read:   resourceAsmAppSyncGraphQLApiRead,
		Update: resourceAsmAppSyncGraphQLApiUpdate,
		Delete: resourceAsmAppSyncGraphQLApiDelete,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"additional_authentication_providers": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_type": {
							Type:     schema.TypeString,
							Required: true,
						},
						"lambda_authorizer_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authorizer_result_ttl_seconds": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  0,
									},
									"authorizer_uri": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"identity_validation_expression": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"openid_connect_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auth_ttl": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"client_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"iat_ttl": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"issuer": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"user_pool_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"app_id_client_regex": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"aws_region": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"user_pool_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"api_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "MERGED",
			},
			"authentication_type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enhanced_metrics_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_source_level_metrics_behavior": {
							Type:     schema.TypeString,
							Required: true,
						},
						"operation_level_metrics_config": {
							Type:     schema.TypeString,
							Required: true,
						},
						"resolver_level_metrics_behavior": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"introspection_config": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "ENABLED",
			},
			"lambda_authorizer_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authorizer_result_ttl_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
						},
						"authorizer_uri": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"identity_validation_expression": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"log_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cloudwatch_logs_role_arn": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"exclude_verbose_content": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"field_log_level": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"merged_api_execution_role_arn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"openid_connect_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"client_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"iat_ttl": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"issuer": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"owner_contact": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"query_depth_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"resolver_count_limit": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"tags": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"user_pool_config": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"app_id_client_regex": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"aws_region": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"default_action": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"user_pool_id": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"visibility": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GLOBAL",
				ForceNew: true,
			},
			"xray_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"schema": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceAsmAppSyncGraphQLApiCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.CreateGraphqlApiInput{
		AdditionalAuthenticationProviders: expandAdditionalAuthenticationProviders(d.Get("additional_authentication_providers").([]interface{})),
		ApiType:                           aws.String(d.Get("api_type").(string)),
		AuthenticationType:                aws.String(d.Get("authentication_type").(string)),
		EnhancedMetricsConfig:             expandEnhancedMetricsConfig(d.Get("enhanced_metrics_config").([]interface{})),
		IntrospectionConfig:               aws.String(d.Get("introspection_config").(string)),
		LambdaAuthorizerConfig:            expandLambdaAuthorizerConfig(d.Get("lambda_authorizer_config").([]interface{})),
		LogConfig:                         expandLogConfig(d.Get("log_config").([]interface{})),
		MergedApiExecutionRoleArn:         aws.String(d.Get("merged_api_execution_role_arn").(string)),
		Name:                              aws.String(d.Get("name").(string)),
		OpenIDConnectConfig:               expandOpenIDConnectConfig(d.Get("openid_connect_config").([]interface{})),
		OwnerContact:                      aws.String(d.Get("owner_contact").(string)),
		QueryDepthLimit:                   aws.Int64(int64(d.Get("query_depth_limit").(int))),
		ResolverCountLimit:                aws.Int64(int64(d.Get("resolver_count_limit").(int))),
		Tags:                              expandTags(d.Get("tags").(map[string]interface{})),
		UserPoolConfig:                    expandUserPoolConfig(d.Get("user_pool_config").([]interface{})),
		Visibility:                        aws.String(d.Get("visibility").(string)),
		XrayEnabled:                       aws.Bool(d.Get("xray_enabled").(bool)),
	}

	result, err := client.AppSync.CreateGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId(aws.StringValue(result.GraphqlApi.ApiId))

	return resourceAsmAppSyncGraphQLApiRead(d, meta)
}

func resourceAsmAppSyncGraphQLApiRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.GetGraphqlApiInput{
		ApiId: aws.String(d.Id()),
	}
	result, err := client.AppSync.GetGraphqlApi(input)
	if err != nil {
		return err
	}

	d.Set("additional_authentication_providers", flattenAdditionalAuthenticationProviders(result.GraphqlApi.AdditionalAuthenticationProviders))
	d.Set("api_type", result.GraphqlApi.ApiType)
	d.Set("authentication_type", result.GraphqlApi.AuthenticationType)
	d.Set("enhanced_metrics_config", flattenEnhancedMetricsConfig(result.GraphqlApi.EnhancedMetricsConfig))
	d.Set("introspection_config", result.GraphqlApi.IntrospectionConfig)
	d.Set("lambda_authorizer_config", flattenLambdaAuthorizerConfig(result.GraphqlApi.LambdaAuthorizerConfig))
	d.Set("log_config", flattenLogConfig(result.GraphqlApi.LogConfig))
	d.Set("merged_api_execution_role_arn", result.GraphqlApi.MergedApiExecutionRoleArn)
	d.Set("name", result.GraphqlApi.Name)
	d.Set("open_id_connect_config", flattenOpenIDConnectConfig(result.GraphqlApi.OpenIDConnectConfig))
	d.Set("owner_contact", result.GraphqlApi.OwnerContact)
	d.Set("query_depth_limit", result.GraphqlApi.QueryDepthLimit)
	d.Set("tags", aws.StringValueMap(result.GraphqlApi.Tags))
	d.Set("user_pool_config", flattenUserPoolConfig(result.GraphqlApi.UserPoolConfig))
	d.Set("visibility", result.GraphqlApi.Visibility)
	d.Set("xray_enabled", result.GraphqlApi.XrayEnabled)
	d.Set("arn", aws.StringValue(result.GraphqlApi.Arn))

	return nil
}

func resourceAsmAppSyncGraphQLApiUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.UpdateGraphqlApiInput{
		ApiId:              aws.String(d.Id()),
		AuthenticationType: aws.String(d.Get("authentication_type").(string)),
		Name:               aws.String(d.Get("name").(string)),
	}

	if d.HasChange("additional_authentication_providers") {
		input.AdditionalAuthenticationProviders = expandAdditionalAuthenticationProviders(d.Get("additional_authentication_providers").([]interface{}))
	}

	if d.HasChange("enhanced_metrics_config") {
		input.EnhancedMetricsConfig = expandEnhancedMetricsConfig(d.Get("enhanced_metrics_config").([]interface{}))
	}

	if d.HasChange("introspection_config") {
		input.IntrospectionConfig = aws.String(d.Get("introspection_config").(string))
	}

	if d.HasChange("lambda_authorizer_config") {
		input.LambdaAuthorizerConfig = expandLambdaAuthorizerConfig(d.Get("lambda_authorizer_config").([]interface{}))
	}

	if d.HasChange("log_config") {
		input.LogConfig = expandLogConfig(d.Get("log_config").([]interface{}))
	}

	if d.HasChange("merged_api_execution_role_arn") {
		input.MergedApiExecutionRoleArn = aws.String(d.Get("merged_api_execution_role_arn").(string))
	}

	if d.HasChange("open_id_connect_config") {
		input.OpenIDConnectConfig = expandOpenIDConnectConfig(d.Get("open_id_connect_config").([]interface{}))
	}

	if d.HasChange("owner_contact") {
		input.OwnerContact = aws.String(d.Get("owner_contact").(string))
	}

	if d.HasChange("query_depth_limit") {
		input.QueryDepthLimit = aws.Int64(int64(d.Get("query_depth_limit").(int)))
	}

	if d.HasChange("resolver_count_limit") {
		input.ResolverCountLimit = aws.Int64(int64(d.Get("resolver_count_limit").(int)))
	}

	if d.HasChange("user_pool_config") {
		input.UserPoolConfig = expandUserPoolConfig(d.Get("user_pool_config").([]interface{}))
	}

	if d.HasChange("xray_enabled") {
		input.XrayEnabled = aws.Bool(d.Get("xray_enabled").(bool))
	}

	if d.HasChange("tags") {
		err := updateTags(d, meta)
		if err != nil {
			return err
		}
	}

	_, err := client.AppSync.UpdateGraphqlApi(input)
	if err != nil {
		return err
	}

	return resourceAsmAppSyncGraphQLApiRead(d, meta)
}

func resourceAsmAppSyncGraphQLApiDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	input := &appsync.DeleteGraphqlApiInput{
		ApiId: aws.String(d.Id()),
	}

	_, err := client.AppSync.DeleteGraphqlApi(input)
	if err != nil {
		return err
	}

	d.SetId("")
	return nil
}

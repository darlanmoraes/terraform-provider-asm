package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/appsync"
)

func expandAdditionalAuthenticationProviders(v []interface{}) []*appsync.AdditionalAuthenticationProvider {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	var result []*appsync.AdditionalAuthenticationProvider
	for _, provider := range v {
		config := provider.(map[string]interface{})
		result = append(result, &appsync.AdditionalAuthenticationProvider{
			AuthenticationType:     aws.String(config["authentication_type"].(string)),
			LambdaAuthorizerConfig: expandLambdaAuthorizerConfig(config["lambda_authorizer_config"].([]interface{})),
			OpenIDConnectConfig:    expandOpenIDConnectConfig(config["openid_connect_config"].([]interface{})),
			UserPoolConfig:         expandCognitoUserPoolConfig(config["user_pool_config"].([]interface{})),
		})
	}
	return result
}

func expandEnhancedMetricsConfig(v []interface{}) *appsync.EnhancedMetricsConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.EnhancedMetricsConfig{
		DataSourceLevelMetricsBehavior: aws.String(config["data_source_level_metrics_behavior"].(string)),
		OperationLevelMetricsConfig:    aws.String(config["operation_level_metrics_config"].(string)),
		ResolverLevelMetricsBehavior:   aws.String(config["resolver_level_metrics_behavior"].(string)),
	}
}

func expandLambdaAuthorizerConfig(v []interface{}) *appsync.LambdaAuthorizerConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.LambdaAuthorizerConfig{
		AuthorizerResultTtlInSeconds: aws.Int64(int64(config["authorizer_result_ttl_seconds"].(int))),
		AuthorizerUri:                aws.String(config["authorizer_uri"].(string)),
		IdentityValidationExpression: aws.String(config["identity_validation_expression"].(string)),
	}
}

func expandLogConfig(v []interface{}) *appsync.LogConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.LogConfig{
		CloudWatchLogsRoleArn: aws.String(config["cloudwatch_logs_role_arn"].(string)),
		ExcludeVerboseContent: aws.Bool(config["exclude_verbose_content"].(bool)),
		FieldLogLevel:         aws.String(config["field_log_level"].(string)),
	}
}

func expandOpenIDConnectConfig(v []interface{}) *appsync.OpenIDConnectConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.OpenIDConnectConfig{
		AuthTTL:  aws.Int64(int64(config["auth_ttl"].(int))),
		ClientId: aws.String(config["client_id"].(string)),
		IatTTL:   aws.Int64(int64(config["iat_ttl"].(int))),
		Issuer:   aws.String(config["issuer"].(string)),
	}
}

func expandTags(v map[string]interface{}) map[string]*string {
	result := make(map[string]*string)
	for k, v := range v {
		result[k] = aws.String(v.(string))
	}
	return result
}

func expandUserPoolConfig(v []interface{}) *appsync.UserPoolConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.UserPoolConfig{
		AppIdClientRegex: aws.String(config["app_id_client_regex"].(string)),
		AwsRegion:        aws.String(config["aws_region"].(string)),
		DefaultAction:    aws.String(config["default_action"].(string)),
		UserPoolId:       aws.String(config["user_pool_id"].(string)),
	}
}

func expandCognitoUserPoolConfig(v []interface{}) *appsync.CognitoUserPoolConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.CognitoUserPoolConfig{
		AppIdClientRegex: aws.String(config["app_id_client_regex"].(string)),
		AwsRegion:        aws.String(config["aws_region"].(string)),
		UserPoolId:       aws.String(config["user_pool_id"].(string)),
	}
}

func expandSourceApiAssociationConfig(v []interface{}) *appsync.SourceApiAssociationConfig {
	if len(v) == 0 || v[0] == nil {
		return nil
	}

	config := v[0].(map[string]interface{})
	return &appsync.SourceApiAssociationConfig{
		MergeType: aws.String(config["merge_type"].(string)),
	}
}

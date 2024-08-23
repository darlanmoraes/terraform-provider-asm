package main

import "github.com/aws/aws-sdk-go/service/appsync"

func flattenLogConfig(v *appsync.LogConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"cloudwatch_logs_role_arn": v.CloudWatchLogsRoleArn,
		"exclude_verbose_content":  v.ExcludeVerboseContent,
		"field_log_level":          v.FieldLogLevel,
	}

	return []interface{}{result}
}

func flattenCognitoUserPoolConfig(v *appsync.CognitoUserPoolConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"aws_region":   v.AwsRegion,
		"user_pool_id": v.UserPoolId,
	}

	if v.AppIdClientRegex != nil {
		result["app_id_client_regex"] = v.AppIdClientRegex
	}

	return []interface{}{result}
}

func flattenUserPoolConfig(v *appsync.UserPoolConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"aws_region":     v.AwsRegion,
		"default_action": v.DefaultAction,
		"user_pool_id":   v.UserPoolId,
	}

	if v.AppIdClientRegex != nil {
		result["app_id_client_regex"] = v.AppIdClientRegex
	}

	return []interface{}{result}
}

func flattenOpenIDConnectConfig(v *appsync.OpenIDConnectConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"auth_ttl":  v.AuthTTL,
		"client_id": v.ClientId,
		"iat_ttl":   v.IatTTL,
		"issuer":    v.Issuer,
	}

	return []interface{}{result}
}

func flattenAdditionalAuthenticationProviders(v []*appsync.AdditionalAuthenticationProvider) []interface{} {
	if len(v) == 0 {
		return []interface{}{}
	}

	var result []interface{}
	for _, config := range v {
		result = append(result, map[string]interface{}{
			"authentication_type":      config.AuthenticationType,
			"lambda_authorizer_config": flattenLambdaAuthorizerConfig(config.LambdaAuthorizerConfig),
			"open_id_connect_config":   flattenOpenIDConnectConfig(config.OpenIDConnectConfig),
			"user_pool_config":         flattenCognitoUserPoolConfig(config.UserPoolConfig),
		})
	}
	return result
}

func flattenEnhancedMetricsConfig(v *appsync.EnhancedMetricsConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"data_source_level_metrics_behavior": v.DataSourceLevelMetricsBehavior,
		"operation_level_metrics_config":     v.OperationLevelMetricsConfig,
		"resolver_level_metrics_behavior":    v.ResolverLevelMetricsBehavior,
	}

	return []interface{}{result}
}

func flattenLambdaAuthorizerConfig(v *appsync.LambdaAuthorizerConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"authorizer_result_ttl_in_seconds": v.AuthorizerResultTtlInSeconds,
		"authorizer_uri":                   v.AuthorizerUri,
	}

	if v.IdentityValidationExpression != nil {
		result["identity_validation_expression"] = v.IdentityValidationExpression
	}

	return []interface{}{result}
}

func flattenSourceApiAssociationConfig(v *appsync.SourceApiAssociationConfig) []interface{} {
	if v == nil {
		return []interface{}{}
	}

	result := map[string]interface{}{
		"merge_type": v.MergeType,
	}

	return []interface{}{result}
}

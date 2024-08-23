# Terraform Provider for ASM

The Terraform provider for ASM allows you to manage AppSync Merged resources using Terraform. This document provides an overview of the provider's configuration, resources, and data sources.

## Overview

The `asm` provider is used to interact with ASM services. This documentation will cover:

- Provider Configuration
- Resources

## Provider Configuration

To use the `asm` provider, you need to specify the required configuration in your Terraform configuration file. The provider uses the same access key and secret key from the environment that the `aws` provider uses. Here’s an example of how to configure the provider:

```hcl
terraform {
  required_providers {
    asm = {
      source  = "darlanmoraes/asm"
      version = "1.0.0"
    }
  }
}

provider "asm" {
  region  = "us-west-1"
}
```

## Usage

**UNDER CONSTRUCTION**

```hcl
resource "asm_appsync_graphql_api" "default" {
  name                = "provider-test-1"
  authentication_type = "AMAZON_COGNITO_USER_POOLS"
  api_type            = "MERGED"

  user_pool_config {
    user_pool_id   = "us-east-1_potato"
    aws_region     = "us-east-1"
    default_action = "ALLOW"
  }

  merged_api_execution_role_arn = "arn:aws:iam::123456:role/potato"
  xray_enabled                  = true

  enhanced_metrics_config {
    resolver_level_metrics_behavior    = "PER_RESOLVER_METRICS"
    data_source_level_metrics_behavior = "PER_DATA_SOURCE_METRICS"
    operation_level_metrics_config     = "ENABLED"
  }

  log_config {
    field_log_level          = "NONE"
    cloudwatch_logs_role_arn = "arn:aws:iam::123456:role/potato"
    exclude_verbose_content  = true
  }

  tags = {
    Name  = "provider-test-1"
    Test1 = "test1"
    Test2 = "test2"
    Test3 = "test3"
  }
}

resource "aws_appsync_graphql_api" "default" {
  authentication_type = "AMAZON_COGNITO_USER_POOLS"
  name                = "provider-test-2"

  user_pool_config {
    aws_region     = "us-east-1"
    default_action = "ALLOW"
    user_pool_id   = "us-east-1_potato"
  }

  tags = {
    Name = "example"
    Test = "test"
  }
}

resource "asm_appsync_merged_api_association" "default" {
  description           = "provider-test-2-into-provider-test-1"
  merged_api_identifier = asm_appsync_graphql_api.default.id
  source_api_identifier = aws_appsync_graphql_api.default.id
  source_api_association_config {
    merge_type = "AUTO_MERGE"
  }
}
```

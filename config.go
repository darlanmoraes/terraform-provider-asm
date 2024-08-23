package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/appsync"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type AWSClient struct {
	Session *session.Session
	AppSync *appsync.AppSync
}

func configureFunc(d *schema.ResourceData) (interface{}, error) {
	region := d.Get("region").(string)

	ses, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return nil, err
	}

	return &AWSClient{
		Session: ses,
		AppSync: appsync.New(ses),
	}, nil
}

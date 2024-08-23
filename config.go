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

func updateTags(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*AWSClient)

	oldTags, newTags := d.GetChange("tags")
	oldTagMap := make(map[string]*string)
	for key, value := range oldTags.(map[string]interface{}) {
		val := value.(string)
		oldTagMap[key] = &val
	}

	newTagMap := make(map[string]*string)
	for key, value := range newTags.(map[string]interface{}) {
		val := value.(string)
		newTagMap[key] = &val
	}

	tagsToAdd := make(map[string]*string)
	tagsToRemove := []*string{}

	for key, value := range newTagMap {
		if oldTagMap[key] == nil || *oldTagMap[key] != *value {
			tagsToAdd[key] = value
		}
	}

	for key := range oldTagMap {
		if newTagMap[key] == nil {
			tagsToRemove = append(tagsToRemove, aws.String(key))
		}
	}

	if len(tagsToAdd) > 0 {
		_, err := client.AppSync.TagResource(&appsync.TagResourceInput{
			ResourceArn: aws.String(d.Get("arn").(string)),
			Tags:        tagsToAdd,
		})
		if err != nil {
			return err
		}
	}

	if len(tagsToRemove) > 0 {
		_, err := client.AppSync.UntagResource(&appsync.UntagResourceInput{
			ResourceArn: aws.String(d.Get("arn").(string)),
			TagKeys:     tagsToRemove,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

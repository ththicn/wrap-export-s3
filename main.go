package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/rds"
)

func main() {
	obj, err := parseYAMLFile("/Users/ththicn/src/github.com/ththicn/wrap-export-s3/resources.yml")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(obj)
	tbls := concatenateDatabaseWithTables(obj.In.Database, obj.In.Tables)

	client := createRDSClient()
	client.StartExportTask(
		context.TODO(),
		&rds.StartExportTaskInput{
			ExportOnly:           tbls,
			ExportTaskIdentifier: aws.String("test"),
			IamRoleArn:           aws.String("my-iam-role-arn"),
			KmsKeyId:             aws.String("my-kms-key-arn"),
			S3BucketName:         aws.String(obj.Out.Bucket),
			S3Prefix:             aws.String(obj.Out.Path),
			SourceArn:            aws.String(obj.In.Database),
		},
	)
}

func createRDSClient() *rds.Client {
	return rds.New(rds.Options{
		Region: "us-west-2",
	})
}

func concatenateDatabaseWithTables(database string, tables []string) []string {
	var concatenated []string
	for _, table := range tables {
		concatenated = append(concatenated, database+"."+table)
	}
	return concatenated
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/goccy/go-yaml"
)

func main() {
	readYaml()
	exportToS3()
}

type inResource struct {
	Type     string
	Database string
	Tables   []string
}

type outResource struct {
	Type   string
	Bucket string
	Path   string
}

type resources struct {
	In  inResource
	Out outResource
}

func readYaml() resources {
	f, err := os.Open("resources.yml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var r resources
	if err := yaml.NewDecoder(f).Decode(&r); err != nil {
		log.Fatal(err)
	}

	log.Printf("in: %+v\n", r.In)
	log.Printf("out: %+v\n", r.Out)

	return r
}

// Assuming the use of AWS SDK for Go v2
func exportToS3() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create an Amazon RDS service client
	rdsClient := rds.NewFromConfig(cfg)

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(cfg)
	fmt.Printf("S3 client created: %v\n", s3Client)

	// Assuming resources are already read and available in r variable of type resources
	r := readYaml()

	// Start the export task
	exportTaskInput := &rds.StartExportTaskInput{
		ExportTaskIdentifier: aws.String(fmt.Sprintf("export-to-s3-%s", r.In.Database)),
		SourceArn:            aws.String(fmt.Sprintf("arn:aws:rds:region:account-id:db:%s", r.In.Database)),
		S3BucketName:         &r.Out.Bucket,
		S3Prefix:             &r.Out.Path,
		IamRoleArn:           aws.String("arn:aws:iam::account-id:role/service-role/access-s3-bucket-role"),
		KmsKeyId:             aws.String("arn:aws:kms:region:account-id:key/key-id"),
		ExportOnly:           r.In.Tables,
	}

	_, err = rdsClient.StartExportTask(context.TODO(), exportTaskInput)
	if err != nil {
		log.Fatalf("failed to start export task, %v", err)
	}

	log.Println("Export task started successfully")
}

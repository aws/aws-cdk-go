package awsdynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3"
)

// Properties for importing data from the S3.
//
// Example:
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import s3 "github.com/aws/aws-cdk-go/awscdk"
//
//   var bucket IBucket
//
//
//   app := cdk.NewApp()
//   stack := cdk.NewStack(app, jsii.String("Stack"))
//
//   dynamodb.NewTable(stack, jsii.String("Table"), &TableProps{
//   	PartitionKey: &Attribute{
//   		Name: jsii.String("id"),
//   		Type: dynamodb.AttributeType_STRING,
//   	},
//   	ImportSource: &ImportSourceSpecification{
//   		CompressionType: dynamodb.InputCompressionType_GZIP,
//   		InputFormat: dynamodb.InputFormat_DynamoDBJson(),
//   		Bucket: *Bucket,
//   		KeyPrefix: jsii.String("prefix"),
//   	},
//   })
//
type ImportSourceSpecification struct {
	// The S3 bucket that is being imported from.
	Bucket awss3.IBucket `field:"required" json:"bucket" yaml:"bucket"`
	// The format of the imported data.
	InputFormat InputFormat `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// The account number of the S3 bucket that is being imported from.
	// Default: - no value.
	//
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// The compression type of the imported data.
	// Default: InputCompressionType.NONE
	//
	CompressionType InputCompressionType `field:"optional" json:"compressionType" yaml:"compressionType"`
	// The key prefix shared by all S3 Objects that are being imported.
	// Default: - no value.
	//
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}


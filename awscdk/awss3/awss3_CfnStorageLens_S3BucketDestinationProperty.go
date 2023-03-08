package awss3


// This resource contains the details of the bucket where the Amazon S3 Storage Lens metrics export will be placed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var sses3 interface{}
//
//   s3BucketDestinationProperty := &s3BucketDestinationProperty{
//   	accountId: jsii.String("accountId"),
//   	arn: jsii.String("arn"),
//   	format: jsii.String("format"),
//   	outputSchemaVersion: jsii.String("outputSchemaVersion"),
//
//   	// the properties below are optional
//   	encryption: &encryptionProperty{
//   		ssekms: &sSEKMSProperty{
//   			keyId: jsii.String("keyId"),
//   		},
//   		sses3: sses3,
//   	},
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnStorageLens_S3BucketDestinationProperty struct {
	// This property contains the details of the AWS account ID of the S3 Storage Lens export bucket destination.
	AccountId *string `field:"required" json:"accountId" yaml:"accountId"`
	// This property contains the details of the ARN of the bucket destination of the S3 Storage Lens export.
	Arn *string `field:"required" json:"arn" yaml:"arn"`
	// This property contains the details of the format of the S3 Storage Lens export bucket destination.
	Format *string `field:"required" json:"format" yaml:"format"`
	// This property contains the details of the output schema version of the S3 Storage Lens export bucket destination.
	OutputSchemaVersion *string `field:"required" json:"outputSchemaVersion" yaml:"outputSchemaVersion"`
	// This property contains the details of the encryption of the bucket destination of the Amazon S3 Storage Lens metrics export.
	Encryption interface{} `field:"optional" json:"encryption" yaml:"encryption"`
	// This property contains the details of the prefix of the bucket destination of the S3 Storage Lens export .
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


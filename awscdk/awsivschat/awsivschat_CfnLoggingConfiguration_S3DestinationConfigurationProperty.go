package awsivschat


// The S3DestinationConfiguration property type specifies an S3 location where chat logs will be stored.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3DestinationConfigurationProperty := &s3DestinationConfigurationProperty{
//   	bucketName: jsii.String("bucketName"),
//   }
//
type CfnLoggingConfiguration_S3DestinationConfigurationProperty struct {
	// Name of the Amazon S3 bucket where chat activity will be logged.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}


package awsivs


// The S3DestinationConfiguration property type describes an S3 location where recorded videos will be stored.
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
type CfnRecordingConfiguration_S3DestinationConfigurationProperty struct {
	// Location (S3 bucket name) where recorded videos will be stored.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
}


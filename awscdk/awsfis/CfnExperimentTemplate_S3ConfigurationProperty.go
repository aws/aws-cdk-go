package awsfis


// Specifies the configuration for experiment logging to Amazon S3 .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &S3ConfigurationProperty{
//   	BucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	Prefix: jsii.String("prefix"),
//   }
//
type CfnExperimentTemplate_S3ConfigurationProperty struct {
	// The name of the destination bucket.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// The bucket prefix.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


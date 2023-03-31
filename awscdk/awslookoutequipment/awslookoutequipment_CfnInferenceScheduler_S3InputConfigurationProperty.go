package awslookoutequipment


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3InputConfigurationProperty := &s3InputConfigurationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnInferenceScheduler_S3InputConfigurationProperty struct {
	// `CfnInferenceScheduler.S3InputConfigurationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnInferenceScheduler.S3InputConfigurationProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


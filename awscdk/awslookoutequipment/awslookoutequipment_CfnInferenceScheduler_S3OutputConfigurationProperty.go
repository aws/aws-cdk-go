package awslookoutequipment


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3OutputConfigurationProperty := &s3OutputConfigurationProperty{
//   	bucket: jsii.String("bucket"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnInferenceScheduler_S3OutputConfigurationProperty struct {
	// `CfnInferenceScheduler.S3OutputConfigurationProperty.Bucket`.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// `CfnInferenceScheduler.S3OutputConfigurationProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


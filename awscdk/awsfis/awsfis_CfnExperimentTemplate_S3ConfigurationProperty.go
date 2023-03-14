package awsfis


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3ConfigurationProperty := &s3ConfigurationProperty{
//   	bucketName: jsii.String("bucketName"),
//
//   	// the properties below are optional
//   	prefix: jsii.String("prefix"),
//   }
//
type CfnExperimentTemplate_S3ConfigurationProperty struct {
	// `CfnExperimentTemplate.S3ConfigurationProperty.BucketName`.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// `CfnExperimentTemplate.S3ConfigurationProperty.Prefix`.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}


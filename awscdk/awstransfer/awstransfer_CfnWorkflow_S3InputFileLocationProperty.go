package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3InputFileLocationProperty := &s3InputFileLocationProperty{
//   	bucket: jsii.String("bucket"),
//   	key: jsii.String("key"),
//   }
//
type CfnWorkflow_S3InputFileLocationProperty struct {
	// `CfnWorkflow.S3InputFileLocationProperty.Bucket`.
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// `CfnWorkflow.S3InputFileLocationProperty.Key`.
	Key *string `field:"optional" json:"key" yaml:"key"`
}


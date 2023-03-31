package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TagProperty := &s3TagProperty{
//   	key: jsii.String("key"),
//   	value: jsii.String("value"),
//   }
//
type CfnWorkflow_S3TagProperty struct {
	// `CfnWorkflow.S3TagProperty.Key`.
	Key *string `field:"required" json:"key" yaml:"key"`
	// `CfnWorkflow.S3TagProperty.Value`.
	Value *string `field:"required" json:"value" yaml:"value"`
}


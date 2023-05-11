package awstransfer


// Specifies the key-value pair that are assigned to a file during the execution of a Tagging step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3TagProperty := &S3TagProperty{
//   	Key: jsii.String("key"),
//   	Value: jsii.String("value"),
//   }
//
type CfnWorkflow_S3TagProperty struct {
	// The name assigned to the tag that you create.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value that corresponds to the key.
	Value *string `field:"required" json:"value" yaml:"value"`
}


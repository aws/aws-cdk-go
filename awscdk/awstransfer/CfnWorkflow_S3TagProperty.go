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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3tag.html
//
type CfnWorkflow_S3TagProperty struct {
	// The name assigned to the tag that you create.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3tag.html#cfn-transfer-workflow-s3tag-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value that corresponds to the key.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3tag.html#cfn-transfer-workflow-s3tag-value
	//
	Value *string `field:"required" json:"value" yaml:"value"`
}


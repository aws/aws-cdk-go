package awstransfer


// Specifies the details for the Amazon S3 location for an input file to a workflow.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3InputFileLocationProperty := &S3InputFileLocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3inputfilelocation.html
//
type CfnWorkflow_S3InputFileLocationProperty struct {
	// Specifies the S3 bucket for the customer input file.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3inputfilelocation.html#cfn-transfer-workflow-s3inputfilelocation-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The name assigned to the file when it was created in Amazon S3.
	//
	// You use the object key to retrieve the object.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-s3inputfilelocation.html#cfn-transfer-workflow-s3inputfilelocation-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}


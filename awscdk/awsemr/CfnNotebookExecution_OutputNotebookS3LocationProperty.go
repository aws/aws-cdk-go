package awsemr


// The Amazon S3 location for the notebook execution output.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   outputNotebookS3LocationProperty := &OutputNotebookS3LocationProperty{
//   	Bucket: jsii.String("bucket"),
//   	Key: jsii.String("key"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-notebookexecution-outputnotebooks3location.html
//
type CfnNotebookExecution_OutputNotebookS3LocationProperty struct {
	// The Amazon S3 bucket that stores the notebook execution output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-notebookexecution-outputnotebooks3location.html#cfn-emr-notebookexecution-outputnotebooks3location-bucket
	//
	Bucket *string `field:"optional" json:"bucket" yaml:"bucket"`
	// The key to the Amazon S3 location that stores the notebook execution output.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-notebookexecution-outputnotebooks3location.html#cfn-emr-notebookexecution-outputnotebooks3location-key
	//
	Key *string `field:"optional" json:"key" yaml:"key"`
}


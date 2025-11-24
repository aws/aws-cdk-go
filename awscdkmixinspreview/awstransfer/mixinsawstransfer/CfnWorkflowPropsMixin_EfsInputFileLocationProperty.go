package mixinsawstransfer


// Specifies the Amazon EFS identifier and the path for the file being used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   efsInputFileLocationProperty := &EfsInputFileLocationProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	Path: jsii.String("path"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-efsinputfilelocation.html
//
type CfnWorkflowPropsMixin_EfsInputFileLocationProperty struct {
	// The identifier of the file system, assigned by Amazon EFS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-efsinputfilelocation.html#cfn-transfer-workflow-efsinputfilelocation-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The pathname for the folder being used by a workflow.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-efsinputfilelocation.html#cfn-transfer-workflow-efsinputfilelocation-path
	//
	Path *string `field:"optional" json:"path" yaml:"path"`
}


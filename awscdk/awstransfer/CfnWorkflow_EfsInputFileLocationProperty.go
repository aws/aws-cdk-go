package awstransfer


// Specifies the Amazon EFS identifier and the path for the file being used.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   efsInputFileLocationProperty := &EfsInputFileLocationProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	Path: jsii.String("path"),
//   }
//
type CfnWorkflow_EfsInputFileLocationProperty struct {
	// The identifier of the file system, assigned by Amazon EFS.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The pathname for the folder being used by a workflow.
	Path *string `field:"optional" json:"path" yaml:"path"`
}


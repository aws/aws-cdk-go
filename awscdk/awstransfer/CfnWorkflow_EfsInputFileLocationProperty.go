package awstransfer


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
	// `CfnWorkflow.EfsInputFileLocationProperty.FileSystemId`.
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// `CfnWorkflow.EfsInputFileLocationProperty.Path`.
	Path *string `field:"optional" json:"path" yaml:"path"`
}


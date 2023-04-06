package awstransfer


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inputFileLocationProperty := &InputFileLocationProperty{
//   	EfsFileLocation: &EfsInputFileLocationProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   		Path: jsii.String("path"),
//   	},
//   	S3FileLocation: &S3InputFileLocationProperty{
//   		Bucket: jsii.String("bucket"),
//   		Key: jsii.String("key"),
//   	},
//   }
//
type CfnWorkflow_InputFileLocationProperty struct {
	// `CfnWorkflow.InputFileLocationProperty.EfsFileLocation`.
	EfsFileLocation interface{} `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// `CfnWorkflow.InputFileLocationProperty.S3FileLocation`.
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}


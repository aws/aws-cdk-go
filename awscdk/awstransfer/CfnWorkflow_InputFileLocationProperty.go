package awstransfer


// Specifies the location for the file that's being processed.
//
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
	// Specifies the details for the Amazon Elastic File System (Amazon EFS) file that's being decrypted.
	EfsFileLocation interface{} `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// Specifies the details for the Amazon S3 file that's being copied or decrypted.
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}


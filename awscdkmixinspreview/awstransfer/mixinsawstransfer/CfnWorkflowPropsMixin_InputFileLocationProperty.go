package mixinsawstransfer


// Specifies the location for the file that's being processed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
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
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-inputfilelocation.html
//
type CfnWorkflowPropsMixin_InputFileLocationProperty struct {
	// Specifies the details for the Amazon Elastic File System (Amazon EFS) file that's being decrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-inputfilelocation.html#cfn-transfer-workflow-inputfilelocation-efsfilelocation
	//
	EfsFileLocation interface{} `field:"optional" json:"efsFileLocation" yaml:"efsFileLocation"`
	// Specifies the details for the Amazon S3 file that's being copied or decrypted.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-transfer-workflow-inputfilelocation.html#cfn-transfer-workflow-inputfilelocation-s3filelocation
	//
	S3FileLocation interface{} `field:"optional" json:"s3FileLocation" yaml:"s3FileLocation"`
}


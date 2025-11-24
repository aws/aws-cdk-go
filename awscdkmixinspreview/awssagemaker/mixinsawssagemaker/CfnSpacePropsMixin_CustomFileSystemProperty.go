package mixinsawssagemaker


// A file system, created by you, that you assign to a user profile or space for an Amazon SageMaker AI Domain.
//
// Permitted users can access this file system in Amazon SageMaker AI Studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customFileSystemProperty := &CustomFileSystemProperty{
//   	EfsFileSystem: &EFSFileSystemProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   	},
//   	FSxLustreFileSystem: &FSxLustreFileSystemProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   	},
//   	S3FileSystem: &S3FileSystemProperty{
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customfilesystem.html
//
type CfnSpacePropsMixin_CustomFileSystemProperty struct {
	// A custom file system in Amazon EFS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customfilesystem.html#cfn-sagemaker-space-customfilesystem-efsfilesystem
	//
	EfsFileSystem interface{} `field:"optional" json:"efsFileSystem" yaml:"efsFileSystem"`
	// A custom file system in Amazon FSx for Lustre.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customfilesystem.html#cfn-sagemaker-space-customfilesystem-fsxlustrefilesystem
	//
	FSxLustreFileSystem interface{} `field:"optional" json:"fSxLustreFileSystem" yaml:"fSxLustreFileSystem"`
	// A custom file system in Amazon S3.
	//
	// This is only supported in Amazon SageMaker Unified Studio.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customfilesystem.html#cfn-sagemaker-space-customfilesystem-s3filesystem
	//
	S3FileSystem interface{} `field:"optional" json:"s3FileSystem" yaml:"s3FileSystem"`
}


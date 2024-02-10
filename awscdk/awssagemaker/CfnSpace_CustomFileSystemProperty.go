package awssagemaker


// A file system, created by you, that you assign to a user profile or space for an Amazon SageMaker Domain.
//
// Permitted users can access this file system in Amazon SageMaker Studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customFileSystemProperty := &CustomFileSystemProperty{
//   	EfsFileSystem: &EFSFileSystemProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customfilesystem.html
//
type CfnSpace_CustomFileSystemProperty struct {
	// A custom file system in Amazon EFS.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-customfilesystem.html#cfn-sagemaker-space-customfilesystem-efsfilesystem
	//
	EfsFileSystem interface{} `field:"optional" json:"efsFileSystem" yaml:"efsFileSystem"`
}


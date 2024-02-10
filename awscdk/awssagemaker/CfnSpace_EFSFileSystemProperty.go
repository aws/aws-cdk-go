package awssagemaker


// A file system, created by you in Amazon EFS, that you assign to a user profile or space for an Amazon SageMaker Domain.
//
// Permitted users can access this file system in Amazon SageMaker Studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eFSFileSystemProperty := &EFSFileSystemProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-efsfilesystem.html
//
type CfnSpace_EFSFileSystemProperty struct {
	// The ID of your Amazon EFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-efsfilesystem.html#cfn-sagemaker-space-efsfilesystem-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
}


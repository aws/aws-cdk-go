package previewawssagemakermixins


// The settings for assigning a custom file system to a user profile or space for an Amazon SageMaker AI Domain.
//
// Permitted users can access this file system in Amazon SageMaker AI Studio.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customFileSystemConfigProperty := &CustomFileSystemConfigProperty{
//   	EfsFileSystemConfig: &EFSFileSystemConfigProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   		FileSystemPath: jsii.String("fileSystemPath"),
//   	},
//   	FSxLustreFileSystemConfig: &FSxLustreFileSystemConfigProperty{
//   		FileSystemId: jsii.String("fileSystemId"),
//   		FileSystemPath: jsii.String("fileSystemPath"),
//   	},
//   	S3FileSystemConfig: &S3FileSystemConfigProperty{
//   		MountPath: jsii.String("mountPath"),
//   		S3Uri: jsii.String("s3Uri"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customfilesystemconfig.html
//
type CfnUserProfilePropsMixin_CustomFileSystemConfigProperty struct {
	// The settings for a custom Amazon EFS file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customfilesystemconfig.html#cfn-sagemaker-userprofile-customfilesystemconfig-efsfilesystemconfig
	//
	EfsFileSystemConfig interface{} `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
	// The settings for a custom Amazon FSx for Lustre file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customfilesystemconfig.html#cfn-sagemaker-userprofile-customfilesystemconfig-fsxlustrefilesystemconfig
	//
	FSxLustreFileSystemConfig interface{} `field:"optional" json:"fSxLustreFileSystemConfig" yaml:"fSxLustreFileSystemConfig"`
	// Configuration settings for a custom Amazon S3 file system.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-customfilesystemconfig.html#cfn-sagemaker-userprofile-customfilesystemconfig-s3filesystemconfig
	//
	S3FileSystemConfig interface{} `field:"optional" json:"s3FileSystemConfig" yaml:"s3FileSystemConfig"`
}


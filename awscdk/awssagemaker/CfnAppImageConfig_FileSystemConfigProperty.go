package awssagemaker


// The Amazon Elastic File System storage configuration for a SageMaker image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfigProperty := &FileSystemConfigProperty{
//   	DefaultGid: jsii.Number(123),
//   	DefaultUid: jsii.Number(123),
//   	MountPath: jsii.String("mountPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html
//
type CfnAppImageConfig_FileSystemConfigProperty struct {
	// The default POSIX group ID (GID).
	//
	// If not specified, defaults to `100` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-defaultgid
	//
	DefaultGid *float64 `field:"optional" json:"defaultGid" yaml:"defaultGid"`
	// The default POSIX user ID (UID).
	//
	// If not specified, defaults to `1000` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-defaultuid
	//
	DefaultUid *float64 `field:"optional" json:"defaultUid" yaml:"defaultUid"`
	// The path within the image to mount the user's EFS home directory.
	//
	// The directory should be empty. If not specified, defaults to * /home/sagemaker-user* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-appimageconfig-filesystemconfig.html#cfn-sagemaker-appimageconfig-filesystemconfig-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}


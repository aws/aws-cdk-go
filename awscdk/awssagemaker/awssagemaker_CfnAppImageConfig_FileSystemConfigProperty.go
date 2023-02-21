package awssagemaker


// The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemConfigProperty := &fileSystemConfigProperty{
//   	defaultGid: jsii.Number(123),
//   	defaultUid: jsii.Number(123),
//   	mountPath: jsii.String("mountPath"),
//   }
//
type CfnAppImageConfig_FileSystemConfigProperty struct {
	// The default POSIX group ID (GID).
	//
	// If not specified, defaults to `100` .
	DefaultGid *float64 `field:"optional" json:"defaultGid" yaml:"defaultGid"`
	// The default POSIX user ID (UID).
	//
	// If not specified, defaults to `1000` .
	DefaultUid *float64 `field:"optional" json:"defaultUid" yaml:"defaultUid"`
	// The path within the image to mount the user's EFS home directory.
	//
	// The directory should be empty. If not specified, defaults to * /home/sagemaker-user* .
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
}


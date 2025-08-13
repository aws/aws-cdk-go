package awssagemaker


// Configuration for the custom Amazon S3 file system.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   s3FileSystemConfigProperty := &S3FileSystemConfigProperty{
//   	MountPath: jsii.String("mountPath"),
//   	S3Uri: jsii.String("s3Uri"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-s3filesystemconfig.html
//
type CfnUserProfile_S3FileSystemConfigProperty struct {
	// The file system path where the Amazon S3 storage location will be mounted within the Amazon SageMaker Studio environment.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-s3filesystemconfig.html#cfn-sagemaker-userprofile-s3filesystemconfig-mountpath
	//
	MountPath *string `field:"optional" json:"mountPath" yaml:"mountPath"`
	// The Amazon S3 URI of the S3 file system configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-userprofile-s3filesystemconfig.html#cfn-sagemaker-userprofile-s3filesystemconfig-s3uri
	//
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}


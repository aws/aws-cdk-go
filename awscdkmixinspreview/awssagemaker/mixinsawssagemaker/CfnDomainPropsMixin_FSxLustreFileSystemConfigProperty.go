package mixinsawssagemaker


// The settings for assigning a custom Amazon FSx for Lustre file system to a user profile or space for an Amazon SageMaker Domain.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fSxLustreFileSystemConfigProperty := &FSxLustreFileSystemConfigProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   	FileSystemPath: jsii.String("fileSystemPath"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-fsxlustrefilesystemconfig.html
//
type CfnDomainPropsMixin_FSxLustreFileSystemConfigProperty struct {
	// The globally unique, 17-digit, ID of the file system, assigned by Amazon FSx for Lustre.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-fsxlustrefilesystemconfig.html#cfn-sagemaker-domain-fsxlustrefilesystemconfig-filesystemid
	//
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// The path to the file system directory that is accessible in Amazon SageMaker Studio.
	//
	// Permitted users can access only this directory and below.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-domain-fsxlustrefilesystemconfig.html#cfn-sagemaker-domain-fsxlustrefilesystemconfig-filesystempath
	//
	FileSystemPath *string `field:"optional" json:"fileSystemPath" yaml:"fileSystemPath"`
}


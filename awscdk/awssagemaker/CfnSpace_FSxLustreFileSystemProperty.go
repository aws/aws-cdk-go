package awssagemaker


// A custom file system in Amazon FSx for Lustre.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fSxLustreFileSystemProperty := &FSxLustreFileSystemProperty{
//   	FileSystemId: jsii.String("fileSystemId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-fsxlustrefilesystem.html
//
type CfnSpace_FSxLustreFileSystemProperty struct {
	// Amazon FSx for Lustre file system ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-space-fsxlustrefilesystem.html#cfn-sagemaker-space-fsxlustrefilesystem-filesystemid
	//
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
}


package awsefs


// A reference to a FileSystem resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemReference := &FileSystemReference{
//   	FileSystemArn: jsii.String("fileSystemArn"),
//   	FileSystemId: jsii.String("fileSystemId"),
//   }
//
type FileSystemReference struct {
	// The ARN of the FileSystem resource.
	FileSystemArn *string `field:"required" json:"fileSystemArn" yaml:"fileSystemArn"`
	// The FileSystemId of the FileSystem resource.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
}


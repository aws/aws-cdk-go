package interfacesawss3files


// A reference to a FileSystem resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemReference := &FileSystemReference{
//   	FileSystemArn: jsii.String("fileSystemArn"),
//   }
//
type FileSystemReference struct {
	// The FileSystemArn of the FileSystem resource.
	FileSystemArn *string `field:"required" json:"fileSystemArn" yaml:"fileSystemArn"`
}


package interfacesawss3files


// A reference to a FileSystemPolicy resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fileSystemPolicyReference := &FileSystemPolicyReference{
//   	FileSystemId: jsii.String("fileSystemId"),
//   }
//
type FileSystemPolicyReference struct {
	// The FileSystemId of the FileSystemPolicy resource.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
}


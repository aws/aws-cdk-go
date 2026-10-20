package interfacesawsworkspaces


// A reference to a Directory resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directoryReference := &DirectoryReference{
//   	DirectoryArn: jsii.String("directoryArn"),
//   }
//
type DirectoryReference struct {
	// The Arn of the Directory resource.
	DirectoryArn *string `field:"required" json:"directoryArn" yaml:"directoryArn"`
}


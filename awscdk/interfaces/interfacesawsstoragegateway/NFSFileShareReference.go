package interfacesawsstoragegateway


// A reference to a NFSFileShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   nFSFileShareReference := &NFSFileShareReference{
//   	FileShareArn: jsii.String("fileShareArn"),
//   }
//
type NFSFileShareReference struct {
	// The FileShareARN of the NFSFileShare resource.
	FileShareArn *string `field:"required" json:"fileShareArn" yaml:"fileShareArn"`
}


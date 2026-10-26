package interfacesawsstoragegateway


// A reference to a SMBFileShare resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   sMBFileShareReference := &SMBFileShareReference{
//   	FileShareArn: jsii.String("fileShareArn"),
//   }
//
type SMBFileShareReference struct {
	// The FileShareARN of the SMBFileShare resource.
	FileShareArn *string `field:"required" json:"fileShareArn" yaml:"fileShareArn"`
}


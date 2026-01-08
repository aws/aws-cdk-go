package interfacesawsappmesh


// A reference to a VirtualNode resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   virtualNodeReference := &VirtualNodeReference{
//   	VirtualNodeArn: jsii.String("virtualNodeArn"),
//   }
//
type VirtualNodeReference struct {
	// The Arn of the VirtualNode resource.
	VirtualNodeArn *string `field:"required" json:"virtualNodeArn" yaml:"virtualNodeArn"`
}


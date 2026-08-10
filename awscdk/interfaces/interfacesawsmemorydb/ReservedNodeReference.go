package interfacesawsmemorydb


// A reference to a ReservedNode resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   reservedNodeReference := &ReservedNodeReference{
//   	ReservedNodeArn: jsii.String("reservedNodeArn"),
//   }
//
type ReservedNodeReference struct {
	// The Arn of the ReservedNode resource.
	ReservedNodeArn *string `field:"required" json:"reservedNodeArn" yaml:"reservedNodeArn"`
}


package interfacesawsram


// A reference to a ResourceAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   resourceAssociationReference := &ResourceAssociationReference{
//   	ResourceArn: jsii.String("resourceArn"),
//   	ResourceShareArn: jsii.String("resourceShareArn"),
//   }
//
type ResourceAssociationReference struct {
	// The ResourceArn of the ResourceAssociation resource.
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
	// The ResourceShareArn of the ResourceAssociation resource.
	ResourceShareArn *string `field:"required" json:"resourceShareArn" yaml:"resourceShareArn"`
}


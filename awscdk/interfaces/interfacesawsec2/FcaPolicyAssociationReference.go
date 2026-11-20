package interfacesawsec2


// A reference to a FcaPolicyAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   fcaPolicyAssociationReference := &FcaPolicyAssociationReference{
//   	FcaPolicyAssociationId: jsii.String("fcaPolicyAssociationId"),
//   }
//
type FcaPolicyAssociationReference struct {
	// The FcaPolicyAssociationId of the FcaPolicyAssociation resource.
	FcaPolicyAssociationId *string `field:"required" json:"fcaPolicyAssociationId" yaml:"fcaPolicyAssociationId"`
}


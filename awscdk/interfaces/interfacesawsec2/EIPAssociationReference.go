package interfacesawsec2


// A reference to a EIPAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   eIPAssociationReference := &EIPAssociationReference{
//   	EipAssociationId: jsii.String("eipAssociationId"),
//   }
//
type EIPAssociationReference struct {
	// The Id of the EIPAssociation resource.
	EipAssociationId *string `field:"required" json:"eipAssociationId" yaml:"eipAssociationId"`
}


package interfacesawsdirectconnect


// A reference to a DirectConnectGatewayAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   directConnectGatewayAssociationReference := &DirectConnectGatewayAssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   }
//
type DirectConnectGatewayAssociationReference struct {
	// The AssociationId of the DirectConnectGatewayAssociation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
}


package interfacesawsec2


// A reference to a ClientVpnTargetNetworkAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientVpnTargetNetworkAssociationReference := &ClientVpnTargetNetworkAssociationReference{
//   	AssociationId: jsii.String("associationId"),
//   	ClientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   }
//
type ClientVpnTargetNetworkAssociationReference struct {
	// The AssociationId of the ClientVpnTargetNetworkAssociation resource.
	AssociationId *string `field:"required" json:"associationId" yaml:"associationId"`
	// The ClientVpnEndpointId of the ClientVpnTargetNetworkAssociation resource.
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
}


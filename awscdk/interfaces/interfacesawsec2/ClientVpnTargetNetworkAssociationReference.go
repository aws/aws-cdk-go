package interfacesawsec2


// A reference to a ClientVpnTargetNetworkAssociation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   clientVpnTargetNetworkAssociationReference := &ClientVpnTargetNetworkAssociationReference{
//   	ClientVpnTargetNetworkAssociationId: jsii.String("clientVpnTargetNetworkAssociationId"),
//   }
//
type ClientVpnTargetNetworkAssociationReference struct {
	// The Id of the ClientVpnTargetNetworkAssociation resource.
	ClientVpnTargetNetworkAssociationId *string `field:"required" json:"clientVpnTargetNetworkAssociationId" yaml:"clientVpnTargetNetworkAssociationId"`
}


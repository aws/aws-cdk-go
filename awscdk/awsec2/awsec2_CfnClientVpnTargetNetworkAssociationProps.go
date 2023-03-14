package awsec2


// Properties for defining a `CfnClientVpnTargetNetworkAssociation`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnClientVpnTargetNetworkAssociationProps := &cfnClientVpnTargetNetworkAssociationProps{
//   	clientVpnEndpointId: jsii.String("clientVpnEndpointId"),
//   	subnetId: jsii.String("subnetId"),
//   }
//
type CfnClientVpnTargetNetworkAssociationProps struct {
	// The ID of the Client VPN endpoint.
	ClientVpnEndpointId *string `field:"required" json:"clientVpnEndpointId" yaml:"clientVpnEndpointId"`
	// The ID of the subnet to associate with the Client VPN endpoint.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


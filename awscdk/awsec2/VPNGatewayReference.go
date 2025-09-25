package awsec2


// A reference to a VPNGateway resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPNGatewayReference := &VPNGatewayReference{
//   	VpnGatewayId: jsii.String("vpnGatewayId"),
//   }
//
type VPNGatewayReference struct {
	// The VPNGatewayId of the VPNGateway resource.
	VpnGatewayId *string `field:"required" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}


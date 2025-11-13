package interfacesawsec2


// A reference to a VPNGatewayRoutePropagation resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPNGatewayRoutePropagationReference := &VPNGatewayRoutePropagationReference{
//   	VpnGatewayRoutePropagationId: jsii.String("vpnGatewayRoutePropagationId"),
//   }
//
type VPNGatewayRoutePropagationReference struct {
	// The Id of the VPNGatewayRoutePropagation resource.
	VpnGatewayRoutePropagationId *string `field:"required" json:"vpnGatewayRoutePropagationId" yaml:"vpnGatewayRoutePropagationId"`
}


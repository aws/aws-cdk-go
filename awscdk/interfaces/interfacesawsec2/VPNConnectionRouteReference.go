package interfacesawsec2


// A reference to a VPNConnectionRoute resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPNConnectionRouteReference := &VPNConnectionRouteReference{
//   	DestinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	VpnConnectionId: jsii.String("vpnConnectionId"),
//   }
//
type VPNConnectionRouteReference struct {
	// The DestinationCidrBlock of the VPNConnectionRoute resource.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The VpnConnectionId of the VPNConnectionRoute resource.
	VpnConnectionId *string `field:"required" json:"vpnConnectionId" yaml:"vpnConnectionId"`
}


package awsec2


// A reference to a VPNConnection resource.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vPNConnectionReference := &VPNConnectionReference{
//   	VpnConnectionId: jsii.String("vpnConnectionId"),
//   }
//
type VPNConnectionReference struct {
	// The VpnConnectionId of the VPNConnection resource.
	VpnConnectionId *string `field:"required" json:"vpnConnectionId" yaml:"vpnConnectionId"`
}


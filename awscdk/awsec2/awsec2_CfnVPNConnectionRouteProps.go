package awsec2


// Properties for defining a `CfnVPNConnectionRoute`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnVPNConnectionRouteProps := &cfnVPNConnectionRouteProps{
//   	destinationCidrBlock: jsii.String("destinationCidrBlock"),
//   	vpnConnectionId: jsii.String("vpnConnectionId"),
//   }
//
type CfnVPNConnectionRouteProps struct {
	// The CIDR block associated with the local subnet of the customer network.
	DestinationCidrBlock *string `field:"required" json:"destinationCidrBlock" yaml:"destinationCidrBlock"`
	// The ID of the VPN connection.
	VpnConnectionId *string `field:"required" json:"vpnConnectionId" yaml:"vpnConnectionId"`
}


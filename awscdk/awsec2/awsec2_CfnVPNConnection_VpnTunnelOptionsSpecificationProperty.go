package awsec2


// The tunnel options for a single VPN tunnel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   vpnTunnelOptionsSpecificationProperty := &vpnTunnelOptionsSpecificationProperty{
//   	preSharedKey: jsii.String("preSharedKey"),
//   	tunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   }
//
type CfnVPNConnection_VpnTunnelOptionsSpecificationProperty struct {
	// The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.
	//
	// Constraints: Allowed characters are alphanumeric characters, periods (.), and underscores (_). Must be between 8 and 64 characters in length and cannot start with zero (0).
	PreSharedKey *string `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
	// The range of inside IP addresses for the tunnel.
	//
	// Any specified CIDR blocks must be unique across all VPN connections that use the same virtual private gateway.
	//
	// Constraints: A size /30 CIDR block from the `169.254.0.0/16` range. The following CIDR blocks are reserved and cannot be used:
	//
	// - `169.254.0.0/30`
	// - `169.254.1.0/30`
	// - `169.254.2.0/30`
	// - `169.254.3.0/30`
	// - `169.254.4.0/30`
	// - `169.254.5.0/30`
	// - `169.254.169.252/30`
	TunnelInsideCidr *string `field:"optional" json:"tunnelInsideCidr" yaml:"tunnelInsideCidr"`
}


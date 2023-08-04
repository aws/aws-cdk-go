package awsec2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import cdk "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var secretValue secretValue
//
//   vpnTunnelOption := &VpnTunnelOption{
//   	PreSharedKey: jsii.String("preSharedKey"),
//   	PreSharedKeySecret: secretValue,
//   	TunnelInsideCidr: jsii.String("tunnelInsideCidr"),
//   }
//
type VpnTunnelOption struct {
	// The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.
	//
	// Allowed characters are
	// alphanumeric characters period `.` and underscores `_`. Must be between 8
	// and 64 characters in length and cannot start with zero (0).
	// Default: an Amazon generated pre-shared key.
	//
	// Deprecated: Use `preSharedKeySecret` instead.
	PreSharedKey *string `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
	// The pre-shared key (PSK) to establish initial authentication between the virtual private gateway and customer gateway.
	//
	// Allowed characters are
	// alphanumeric characters period `.` and underscores `_`. Must be between 8
	// and 64 characters in length and cannot start with zero (0).
	// Default: an Amazon generated pre-shared key.
	//
	PreSharedKeySecret awscdk.SecretValue `field:"optional" json:"preSharedKeySecret" yaml:"preSharedKeySecret"`
	// The range of inside IP addresses for the tunnel.
	//
	// Any specified CIDR blocks must be
	// unique across all VPN connections that use the same virtual private gateway.
	// A size /30 CIDR block from the 169.254.0.0/16 range.
	// Default: an Amazon generated inside IP CIDR.
	//
	TunnelInsideCidr *string `field:"optional" json:"tunnelInsideCidr" yaml:"tunnelInsideCidr"`
}


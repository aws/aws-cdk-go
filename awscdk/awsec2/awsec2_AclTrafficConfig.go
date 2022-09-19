package awsec2


// Acl Configuration for traffic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   aclTrafficConfig := &aclTrafficConfig{
//   	protocol: jsii.Number(123),
//
//   	// the properties below are optional
//   	icmp: &aclIcmp{
//   		code: jsii.Number(123),
//   		type: jsii.Number(123),
//   	},
//   	portRange: &aclPortRange{
//   		from: jsii.Number(123),
//   		to: jsii.Number(123),
//   	},
//   }
//
type AclTrafficConfig struct {
	// The protocol number.
	//
	// A value of "-1" means all protocols.
	//
	// If you specify "-1" or a protocol number other than "6" (TCP), "17" (UDP),
	// or "1" (ICMP), traffic on all ports is allowed, regardless of any ports or
	// ICMP types or codes that you specify.
	//
	// If you specify protocol "58" (ICMPv6) and specify an IPv4 CIDR
	// block, traffic for all ICMP types and codes allowed, regardless of any that
	// you specify. If you specify protocol "58" (ICMPv6) and specify an IPv6 CIDR
	// block, you must specify an ICMP type and code.
	Protocol *float64 `field:"required" json:"protocol" yaml:"protocol"`
	// The Internet Control Message Protocol (ICMP) code and type.
	Icmp *AclIcmp `field:"optional" json:"icmp" yaml:"icmp"`
	// The range of port numbers for the UDP/TCP protocol.
	PortRange *AclPortRange `field:"optional" json:"portRange" yaml:"portRange"`
}


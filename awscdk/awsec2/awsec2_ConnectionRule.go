package awsec2


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   connectionRule := &connectionRule{
//   	fromPort: jsii.Number(123),
//
//   	// the properties below are optional
//   	description: jsii.String("description"),
//   	protocol: jsii.String("protocol"),
//   	toPort: jsii.Number(123),
//   }
//
type ConnectionRule struct {
	// Start of port range for the TCP and UDP protocols, or an ICMP type number.
	//
	// If you specify icmp for the IpProtocol property, you can specify
	// -1 as a wildcard (i.e., any ICMP type number).
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Description of this connection.
	//
	// It is applied to both the ingress rule
	// and the egress rule.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The IP protocol name (tcp, udp, icmp) or number (see Protocol Numbers).
	//
	// Use -1 to specify all protocols. If you specify -1, or a protocol number
	// other than tcp, udp, icmp, or 58 (ICMPv6), traffic on all ports is
	// allowed, regardless of any ports you specify. For tcp, udp, and icmp, you
	// must specify a port range. For protocol 58 (ICMPv6), you can optionally
	// specify a port range; if you don't, traffic for all types and codes is
	// allowed.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// End of port range for the TCP and UDP protocols, or an ICMP code.
	//
	// If you specify icmp for the IpProtocol property, you can specify -1 as a
	// wildcard (i.e., any ICMP code).
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}


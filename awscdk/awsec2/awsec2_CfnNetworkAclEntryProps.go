package awsec2


// Properties for defining a `CfnNetworkAclEntry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkAclEntryProps := &cfnNetworkAclEntryProps{
//   	networkAclId: jsii.String("networkAclId"),
//   	protocol: jsii.Number(123),
//   	ruleAction: jsii.String("ruleAction"),
//   	ruleNumber: jsii.Number(123),
//
//   	// the properties below are optional
//   	cidrBlock: jsii.String("cidrBlock"),
//   	egress: jsii.Boolean(false),
//   	icmp: &icmpProperty{
//   		code: jsii.Number(123),
//   		type: jsii.Number(123),
//   	},
//   	ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	portRange: &portRangeProperty{
//   		from: jsii.Number(123),
//   		to: jsii.Number(123),
//   	},
//   }
//
type CfnNetworkAclEntryProps struct {
	// The ID of the ACL for the entry.
	NetworkAclId *string `field:"required" json:"networkAclId" yaml:"networkAclId"`
	// The IP protocol that the rule applies to.
	//
	// You must specify -1 or a protocol number. You can specify -1 for all protocols.
	//
	// > If you specify -1, all ports are opened and the `PortRange` property is ignored.
	Protocol *float64 `field:"required" json:"protocol" yaml:"protocol"`
	// Whether to allow or deny traffic that matches the rule;
	//
	// valid values are "allow" or "deny".
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Rule number to assign to the entry, such as 100.
	//
	// ACL entries are processed in ascending order by rule number. Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule.
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). Requirement is conditional: You must specify the `CidrBlock` or `Ipv6CidrBlock` property.
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Whether this rule applies to egress traffic from the subnet ( `true` ) or ingress traffic to the subnet ( `false` ).
	//
	// By default, AWS CloudFormation specifies `false` .
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// The Internet Control Message Protocol (ICMP) code and type.
	//
	// Requirement is conditional: Required if specifying 1 (ICMP) for the protocol parameter.
	Icmp interface{} `field:"optional" json:"icmp" yaml:"icmp"`
	// The IPv6 network range to allow or deny, in CIDR notation.
	//
	// Requirement is conditional: You must specify the `CidrBlock` or `Ipv6CidrBlock` property.
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The range of port numbers for the UDP/TCP protocol.
	//
	// Conditional required if specifying 6 (TCP) or 17 (UDP) for the protocol parameter.
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
}


package awsec2


// Properties for defining a `CfnNetworkAclEntry`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnNetworkAclEntryProps := &CfnNetworkAclEntryProps{
//   	NetworkAclId: jsii.String("networkAclId"),
//   	Protocol: jsii.Number(123),
//   	RuleAction: jsii.String("ruleAction"),
//   	RuleNumber: jsii.Number(123),
//
//   	// the properties below are optional
//   	CidrBlock: jsii.String("cidrBlock"),
//   	Egress: jsii.Boolean(false),
//   	Icmp: &IcmpProperty{
//   		Code: jsii.Number(123),
//   		Type: jsii.Number(123),
//   	},
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	PortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html
//
type CfnNetworkAclEntryProps struct {
	// The ID of the ACL for the entry.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-networkaclid
	//
	NetworkAclId *string `field:"required" json:"networkAclId" yaml:"networkAclId"`
	// The IP protocol that the rule applies to.
	//
	// You must specify -1 or a protocol number. You can specify -1 for all protocols.
	//
	// > If you specify -1, all ports are opened and the `PortRange` property is ignored.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-protocol
	//
	Protocol *float64 `field:"required" json:"protocol" yaml:"protocol"`
	// Whether to allow or deny traffic that matches the rule;
	//
	// valid values are "allow" or "deny".
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-ruleaction
	//
	RuleAction *string `field:"required" json:"ruleAction" yaml:"ruleAction"`
	// Rule number to assign to the entry, such as 100.
	//
	// ACL entries are processed in ascending order by rule number. Entries can't use the same rule number unless one is an egress rule and the other is an ingress rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-rulenumber
	//
	RuleNumber *float64 `field:"required" json:"ruleNumber" yaml:"ruleNumber"`
	// The IPv4 CIDR range to allow or deny, in CIDR notation (for example, 172.16.0.0/24). You must specify an IPv4 CIDR block or an IPv6 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Whether this rule applies to egress traffic from the subnet ( `true` ) or ingress traffic to the subnet ( `false` ).
	//
	// By default, AWS CloudFormation specifies `false` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-egress
	//
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// The Internet Control Message Protocol (ICMP) code and type.
	//
	// Required if specifying 1 (ICMP) for the protocol parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-icmp
	//
	Icmp interface{} `field:"optional" json:"icmp" yaml:"icmp"`
	// The IPv6 network range to allow or deny, in CIDR notation.
	//
	// You must specify an IPv4 CIDR block or an IPv6 CIDR block.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-ipv6cidrblock
	//
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// The range of port numbers for the UDP/TCP protocol.
	//
	// Required if specifying 6 (TCP) or 17 (UDP) for the protocol parameter.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-networkaclentry.html#cfn-ec2-networkaclentry-portrange
	//
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
}


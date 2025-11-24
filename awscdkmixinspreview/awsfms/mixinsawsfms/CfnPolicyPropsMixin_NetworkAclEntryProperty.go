package mixinsawsfms


// Describes a rule in a network ACL.
//
// Each network ACL has a set of numbered ingress rules and a separate set of numbered egress rules. When determining
// whether a packet should be allowed in or out of a subnet associated with the network ACL, AWS processes the entries in the network ACL according to the rule numbers, in ascending order.
//
// When you manage an individual network ACL, you explicitly specify the rule numbers. When you specify the network ACL rules in a Firewall Manager policy, you provide the rules to run first, in the order that you want them to run, and the rules to run last, in the order that you want them to run. Firewall Manager assigns the rule numbers for you when you save the network ACL policy specification.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   networkAclEntryProperty := &NetworkAclEntryProperty{
//   	CidrBlock: jsii.String("cidrBlock"),
//   	Egress: jsii.Boolean(false),
//   	IcmpTypeCode: &IcmpTypeCodeProperty{
//   		Code: jsii.Number(123),
//   		Type: jsii.Number(123),
//   	},
//   	Ipv6CidrBlock: jsii.String("ipv6CidrBlock"),
//   	PortRange: &PortRangeProperty{
//   		From: jsii.Number(123),
//   		To: jsii.Number(123),
//   	},
//   	Protocol: jsii.String("protocol"),
//   	RuleAction: jsii.String("ruleAction"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html
//
type CfnPolicyPropsMixin_NetworkAclEntryProperty struct {
	// The IPv4 network range to allow or deny, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-cidrblock
	//
	CidrBlock *string `field:"optional" json:"cidrBlock" yaml:"cidrBlock"`
	// Indicates whether the rule is an egress, or outbound, rule (applied to traffic leaving the subnet).
	//
	// If it's not an egress rule, then it's an ingress, or inbound, rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-egress
	//
	Egress interface{} `field:"optional" json:"egress" yaml:"egress"`
	// ICMP protocol: The ICMP type and code.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-icmptypecode
	//
	IcmpTypeCode interface{} `field:"optional" json:"icmpTypeCode" yaml:"icmpTypeCode"`
	// The IPv6 network range to allow or deny, in CIDR notation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-ipv6cidrblock
	//
	Ipv6CidrBlock *string `field:"optional" json:"ipv6CidrBlock" yaml:"ipv6CidrBlock"`
	// TCP or UDP protocols: The range of ports the rule applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-portrange
	//
	PortRange interface{} `field:"optional" json:"portRange" yaml:"portRange"`
	// The protocol number.
	//
	// A value of "-1" means all protocols.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Indicates whether to allow or deny the traffic that matches the rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fms-policy-networkaclentry.html#cfn-fms-policy-networkaclentry-ruleaction
	//
	RuleAction *string `field:"optional" json:"ruleAction" yaml:"ruleAction"`
}


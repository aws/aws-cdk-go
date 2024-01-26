package awsec2


// Adds an inbound (ingress) rule to a security group.
//
// An inbound rule permits instances to receive traffic from the specified IPv4 or IPv6 address range, the IP address ranges that are specified by a prefix list, or the instances that are associated with a source security group. For more information, see [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html) .
//
// You must specify exactly one of the following sources: an IPv4 or IPv6 address range, a prefix list, or a security group. Otherwise, the stack launches successfully, but the rule is not added to the security group.
//
// You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also specify the ICMP/ICMPv6 type and code.
//
// Rule changes are propagated to instances associated with the security group as quickly as possible. However, a small delay might occur.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressProperty := &IngressProperty{
//   	IpProtocol: jsii.String("ipProtocol"),
//
//   	// the properties below are optional
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   	Description: jsii.String("description"),
//   	FromPort: jsii.Number(123),
//   	SourcePrefixListId: jsii.String("sourcePrefixListId"),
//   	SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   	SourceSecurityGroupName: jsii.String("sourceSecurityGroupName"),
//   	SourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html
//
type CfnSecurityGroup_IngressProperty struct {
	// The IP protocol name ( `tcp` , `udp` , `icmp` , `icmpv6` ) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) ).
	//
	// Use `-1` to specify all protocols. When authorizing security group rules, specifying `-1` or a protocol number other than `tcp` , `udp` , `icmp` , or `icmpv6` allows traffic on all ports, regardless of any port range you specify. For `tcp` , `udp` , and `icmp` , you must specify a port range. For `icmpv6` , the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-ipprotocol
	//
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 address range, in CIDR format.
	//
	// You must specify a source security group ( `SourcePrefixListId` or `SourceSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-cidrip
	//
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//
	// You must specify a source security group ( `SourcePrefixListId` or `SourceSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-cidripv6
	//
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// Updates the description of an ingress (inbound) security group rule.
	//
	// You can replace an existing description, or add a description to a rule that did not have one previously.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// If the protocol is TCP or UDP, this is the start of the port range.
	//
	// If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The ID of a prefix list.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-sourceprefixlistid
	//
	SourcePrefixListId *string `field:"optional" json:"sourcePrefixListId" yaml:"sourcePrefixListId"`
	// The ID of the security group.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-sourcesecuritygroupid
	//
	SourceSecurityGroupId *string `field:"optional" json:"sourceSecurityGroupId" yaml:"sourceSecurityGroupId"`
	// [Default VPC] The name of the source security group.
	//
	// You must specify either the security group ID or the security group name. You can't specify the group name in combination with an IP address range. Creates rules that grant full ICMP, UDP, and TCP access.
	//
	// For security groups in a nondefault VPC, you must specify the group ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-sourcesecuritygroupname
	//
	SourceSecurityGroupName *string `field:"optional" json:"sourceSecurityGroupName" yaml:"sourceSecurityGroupName"`
	// [nondefault VPC] The AWS account ID for the source security group, if the source security group is in a different account.
	//
	// You can't specify this property with an IP address range. Creates rules that grant full ICMP, UDP, and TCP access.
	//
	// If you specify `SourceSecurityGroupName` or `SourceSecurityGroupId` and that security group is owned by a different account than the account creating the stack, you must specify the `SourceSecurityGroupOwnerId` ; otherwise, this property is optional.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-sourcesecuritygroupownerid
	//
	SourceSecurityGroupOwnerId *string `field:"optional" json:"sourceSecurityGroupOwnerId" yaml:"sourceSecurityGroupOwnerId"`
	// If the protocol is TCP or UDP, this is the end of the port range.
	//
	// If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-ingress.html#cfn-ec2-securitygroup-ingress-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}


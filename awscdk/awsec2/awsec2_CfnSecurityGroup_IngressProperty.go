package awsec2


// Adds an inbound rule to a security group.
//
// An inbound rule permits instances to receive traffic from the specified IPv4 or IPv6 CIDR address range, or from the instances associated with the specified security group.
//
// You must specify only one of the following properties: `CidrIp` , `CidrIpv6` , `SourcePrefixListId` , `SourceSecurityGroupId` , or `SourceSecurityGroupName` .
//
// You specify a protocol for each rule (for example, TCP). For TCP and UDP, you must also specify a port or port range. For ICMP/ICMPv6, you must also specify the ICMP/ICMPv6 type and code. You can use -1 to mean all types or all codes.
//
// You must specify a source security group ( `SourcePrefixListId` , `SourceSecurityGroupId` , or `SourceSecurityGroupName` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ). If you do not specify one of these parameters, the stack will launch successfully but the rule will not be added to the security group.
//
// Rule changes are propagated to instances within the security group as quickly as possible. However, a small delay might occur.
//
// The EC2 Security Group Rule is an embedded property of the `AWS::EC2::SecurityGroup` type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressProperty := &ingressProperty{
//   	ipProtocol: jsii.String("ipProtocol"),
//
//   	// the properties below are optional
//   	cidrIp: jsii.String("cidrIp"),
//   	cidrIpv6: jsii.String("cidrIpv6"),
//   	description: jsii.String("description"),
//   	fromPort: jsii.Number(123),
//   	sourcePrefixListId: jsii.String("sourcePrefixListId"),
//   	sourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   	sourceSecurityGroupName: jsii.String("sourceSecurityGroupName"),
//   	sourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   	toPort: jsii.Number(123),
//   }
//
type CfnSecurityGroup_IngressProperty struct {
	// The IP protocol name ( `tcp` , `udp` , `icmp` , `icmpv6` ) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) ).
	//
	// [VPC only] Use `-1` to specify all protocols. When authorizing security group rules, specifying `-1` or a protocol number other than `tcp` , `udp` , `icmp` , or `icmpv6` allows traffic on all ports, regardless of any port range you specify. For `tcp` , `udp` , and `icmp` , you must specify a port range. For `icmpv6` , the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 address range, in CIDR format.
	//
	// You must specify a source security group ( `SourcePrefixListId` or `SourceSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//
	// You must specify a source security group ( `SourcePrefixListId` or `SourceSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// Updates the description of an ingress (inbound) security group rule.
	//
	// You can replace an existing description, or add a description to a rule that did not have one previously.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.
	//
	// A value of `-1` indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// [EC2-VPC only] The ID of a prefix list.
	SourcePrefixListId *string `field:"optional" json:"sourcePrefixListId" yaml:"sourcePrefixListId"`
	// The ID of the security group.
	//
	// You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	SourceSecurityGroupId *string `field:"optional" json:"sourceSecurityGroupId" yaml:"sourceSecurityGroupId"`
	// [EC2-Classic, default VPC] The name of the source security group.
	//
	// You can't specify this parameter in combination with an IP address range. Creates rules that grant full ICMP, UDP, and TCP access.
	//
	// You must specify the `GroupName` property or the `GroupId` property. For security groups that are in a VPC, you must use the `GroupId` property.
	SourceSecurityGroupName *string `field:"optional" json:"sourceSecurityGroupName" yaml:"sourceSecurityGroupName"`
	// [nondefault VPC] The AWS account ID for the source security group, if the source security group is in a different account.
	//
	// You can't specify this property with an IP address range. Creates rules that grant full ICMP, UDP, and TCP access.
	//
	// If you specify `SourceSecurityGroupName` or `SourceSecurityGroupId` and that security group is owned by a different account than the account creating the stack, you must specify the `SourceSecurityGroupOwnerId` ; otherwise, this property is optional.
	SourceSecurityGroupOwnerId *string `field:"optional" json:"sourceSecurityGroupOwnerId" yaml:"sourceSecurityGroupOwnerId"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	//
	// A value of `-1` indicates all ICMP/ICMPv6 codes. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}


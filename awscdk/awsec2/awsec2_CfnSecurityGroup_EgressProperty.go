package awsec2


// [EC2-VPC only] Adds the specified egress rules to a security group for use with a VPC.
//
// An outbound rule permits instances to send traffic to the specified destination IPv4 or IPv6 CIDR address ranges, or to the specified destination security groups for the same VPC.
//
// You specify a protocol for each rule (for example, TCP). For the TCP and UDP protocols, you must also specify the destination port or port range. For the ICMP protocol, you must also specify the ICMP type and code. You can use -1 for the type or code to mean all types or all codes.
//
// You must specify only one of the following properties: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
//
// You must specify a destination security group ( `DestinationPrefixListId` or `DestinationSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ). If you do not specify one of these parameters, the stack will launch successfully but the rule will not be added to the security group.
//
// Rule changes are propagated to affected instances as quickly as possible. However, a small delay might occur.
//
// For more information about VPC security group limits, see [Amazon VPC Limits](https://docs.aws.amazon.com/vpc/latest/userguide/amazon-vpc-limits.html) .
//
// Use `SecurityGroup.Ingress` and `SecurityGroup.Egress` only when necessary, typically to allow security groups to reference each other in ingress and egress rules. Otherwise, use the embedded ingress and egress rules of the security group. For more information, see [Amazon EC2 Security Groups](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-network-security.html) .
//
// The EC2 Security Group Rule is an embedded property of the `AWS::EC2::SecurityGroup` type.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressProperty := &egressProperty{
//   	ipProtocol: jsii.String("ipProtocol"),
//
//   	// the properties below are optional
//   	cidrIp: jsii.String("cidrIp"),
//   	cidrIpv6: jsii.String("cidrIpv6"),
//   	description: jsii.String("description"),
//   	destinationPrefixListId: jsii.String("destinationPrefixListId"),
//   	destinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   	fromPort: jsii.Number(123),
//   	toPort: jsii.Number(123),
//   }
//
type CfnSecurityGroup_EgressProperty struct {
	// The IP protocol name ( `tcp` , `udp` , `icmp` , `icmpv6` ) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) ).
	//
	// [VPC only] Use `-1` to specify all protocols. When authorizing security group rules, specifying `-1` or a protocol number other than `tcp` , `udp` , `icmp` , or `icmpv6` allows traffic on all ports, regardless of any port range you specify. For `tcp` , `udp` , and `icmp` , you must specify a port range. For `icmpv6` , the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 address range, in CIDR format.
	//
	// You must specify a destination security group ( `DestinationPrefixListId` or `DestinationSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//
	// You must specify a destination security group ( `DestinationPrefixListId` or `DestinationSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// A description for the security group rule.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	Description *string `field:"optional" json:"description" yaml:"description"`
	// [EC2-VPC only] The prefix list IDs for the destination AWS service.
	//
	// This is the AWS service that you want to access through a VPC endpoint from instances associated with the security group.
	//
	// You must specify a destination security group ( `DestinationPrefixListId` or `DestinationSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// The ID of the destination VPC security group.
	//
	// You must specify a destination security group ( `DestinationPrefixListId` or `DestinationSecurityGroupId` ) or a CIDR range ( `CidrIp` or `CidrIpv6` ).
	DestinationSecurityGroupId *string `field:"optional" json:"destinationSecurityGroupId" yaml:"destinationSecurityGroupId"`
	// The start of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 type number.
	//
	// A value of `-1` indicates all ICMP/ICMPv6 types. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// The end of port range for the TCP and UDP protocols, or an ICMP/ICMPv6 code.
	//
	// A value of `-1` indicates all ICMP/ICMPv6 codes. If you specify all ICMP/ICMPv6 types, you must specify all codes.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}


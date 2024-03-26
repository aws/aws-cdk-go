package awsec2


// Adds the specified outbound (egress) rule to a security group.
//
// An outbound rule permits instances to send traffic to the specified IPv4 or IPv6 address range, the IP address ranges that are specified by a prefix list, or the instances that are associated with a destination security group. For more information, see [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html) .
//
// You must specify exactly one of the following destinations: an IPv4 address range, an IPv6 address range, a prefix list, or a security group.
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
//   egressProperty := &EgressProperty{
//   	IpProtocol: jsii.String("ipProtocol"),
//
//   	// the properties below are optional
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   	Description: jsii.String("description"),
//   	DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   	DestinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   	FromPort: jsii.Number(123),
//   	SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html
//
type CfnSecurityGroup_EgressProperty struct {
	// The IP protocol name ( `tcp` , `udp` , `icmp` , `icmpv6` ) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) ).
	//
	// Use `-1` to specify all protocols. When authorizing security group rules, specifying `-1` or a protocol number other than `tcp` , `udp` , `icmp` , or `icmpv6` allows traffic on all ports, regardless of any port range you specify. For `tcp` , `udp` , and `icmp` , you must specify a port range. For `icmpv6` , the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-ipprotocol
	//
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 address range, in CIDR format.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-cidrip
	//
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-cidripv6
	//
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// A description for the security group rule.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The prefix list IDs for the destination AWS service.
	//
	// This is the AWS service that you want to access through a VPC endpoint from instances associated with the security group.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-destinationprefixlistid
	//
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// The ID of the destination VPC security group.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-destinationsecuritygroupid
	//
	DestinationSecurityGroupId *string `field:"optional" json:"destinationSecurityGroupId" yaml:"destinationSecurityGroupId"`
	// If the protocol is TCP or UDP, this is the start of the port range.
	//
	// If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-sourcesecuritygroupid
	//
	SourceSecurityGroupId *string `field:"optional" json:"sourceSecurityGroupId" yaml:"sourceSecurityGroupId"`
	// If the protocol is TCP or UDP, this is the end of the port range.
	//
	// If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-securitygroup-egress.html#cfn-ec2-securitygroup-egress-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}


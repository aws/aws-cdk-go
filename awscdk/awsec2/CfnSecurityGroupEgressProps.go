package awsec2


// Properties for defining a `CfnSecurityGroupEgress`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnSecurityGroupEgressProps := &CfnSecurityGroupEgressProps{
//   	GroupId: jsii.String("groupId"),
//   	IpProtocol: jsii.String("ipProtocol"),
//
//   	// the properties below are optional
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   	Description: jsii.String("description"),
//   	DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   	DestinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   	FromPort: jsii.Number(123),
//   	ToPort: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html
//
type CfnSecurityGroupEgressProps struct {
	// The ID of the security group.
	//
	// You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-groupid
	//
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The IP protocol name ( `tcp` , `udp` , `icmp` , `icmpv6` ) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml) ).
	//
	// Use `-1` to specify all protocols. When authorizing security group rules, specifying `-1` or a protocol number other than `tcp` , `udp` , `icmp` , or `icmpv6` allows traffic on all ports, regardless of any port range you specify. For `tcp` , `udp` , and `icmp` , you must specify a port range. For `icmpv6` , the port range is optional; if you omit the port range, traffic for all types and codes is allowed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-ipprotocol
	//
	IpProtocol *string `field:"required" json:"ipProtocol" yaml:"ipProtocol"`
	// The IPv4 address range, in CIDR format.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-cidrip
	//
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	//
	// For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *Amazon EC2 User Guide* .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-cidripv6
	//
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// The description of an egress (outbound) security group rule.
	//
	// Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The prefix list IDs for an AWS service.
	//
	// This is the AWS service to access through a VPC endpoint from instances associated with the security group.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-destinationprefixlistid
	//
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// The ID of the security group.
	//
	// You must specify exactly one of the following: `CidrIp` , `CidrIpv6` , `DestinationPrefixListId` , or `DestinationSecurityGroupId` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-destinationsecuritygroupid
	//
	DestinationSecurityGroupId *string `field:"optional" json:"destinationSecurityGroupId" yaml:"destinationSecurityGroupId"`
	// If the protocol is TCP or UDP, this is the start of the port range.
	//
	// If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-fromport
	//
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// If the protocol is TCP or UDP, this is the end of the port range.
	//
	// If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-securitygroupegress.html#cfn-ec2-securitygroupegress-toport
	//
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}


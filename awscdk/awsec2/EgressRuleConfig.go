package awsec2


// Configuration for an egress security group rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   egressRuleConfig := &EgressRuleConfig{
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   	DestinationPrefixListId: jsii.String("destinationPrefixListId"),
//   	DestinationSecurityGroupId: jsii.String("destinationSecurityGroupId"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-securitygroup-egress.html
//
type EgressRuleConfig struct {
	// The IPv4 address range, in CIDR format.
	// Default: - No IPv4 CIDR.
	//
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	// Default: - No IPv6 CIDR.
	//
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// The ID of a destination prefix list.
	// Default: - No destination prefix list.
	//
	DestinationPrefixListId *string `field:"optional" json:"destinationPrefixListId" yaml:"destinationPrefixListId"`
	// The ID of a destination security group.
	// Default: - No destination security group.
	//
	DestinationSecurityGroupId *string `field:"optional" json:"destinationSecurityGroupId" yaml:"destinationSecurityGroupId"`
}


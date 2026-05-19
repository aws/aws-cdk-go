package awsec2


// Configuration for an ingress security group rule.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ingressRuleConfig := &IngressRuleConfig{
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   	SourcePrefixListId: jsii.String("sourcePrefixListId"),
//   	SourceSecurityGroupId: jsii.String("sourceSecurityGroupId"),
//   	SourceSecurityGroupOwnerId: jsii.String("sourceSecurityGroupOwnerId"),
//   }
//
// See: https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ec2-securitygroup-ingress.html
//
type IngressRuleConfig struct {
	// The IPv4 address range, in CIDR format.
	// Default: - No IPv4 CIDR.
	//
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	// Default: - No IPv6 CIDR.
	//
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
	// The ID of a source prefix list.
	// Default: - No source prefix list.
	//
	SourcePrefixListId *string `field:"optional" json:"sourcePrefixListId" yaml:"sourcePrefixListId"`
	// The ID of a source security group.
	// Default: - No source security group.
	//
	SourceSecurityGroupId *string `field:"optional" json:"sourceSecurityGroupId" yaml:"sourceSecurityGroupId"`
	// The AWS account ID of the owner of a source security group.
	// Default: - No source security group owner ID.
	//
	SourceSecurityGroupOwnerId *string `field:"optional" json:"sourceSecurityGroupOwnerId" yaml:"sourceSecurityGroupOwnerId"`
}


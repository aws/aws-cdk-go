package awsec2


// Common configuration properties shared by ingress and egress security group rules.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   ruleConfig := &RuleConfig{
//   	CidrIp: jsii.String("cidrIp"),
//   	CidrIpv6: jsii.String("cidrIpv6"),
//   }
//
type RuleConfig struct {
	// The IPv4 address range, in CIDR format.
	// Default: - No IPv4 CIDR.
	//
	CidrIp *string `field:"optional" json:"cidrIp" yaml:"cidrIp"`
	// The IPv6 address range, in CIDR format.
	// Default: - No IPv6 CIDR.
	//
	CidrIpv6 *string `field:"optional" json:"cidrIpv6" yaml:"cidrIpv6"`
}


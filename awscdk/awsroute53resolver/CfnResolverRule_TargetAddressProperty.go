package awsroute53resolver


// In a [CreateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html) request, an array of the IPs that you want to forward DNS queries to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   targetAddressProperty := &TargetAddressProperty{
//   	Ip: jsii.String("ip"),
//   	Ipv6: jsii.String("ipv6"),
//   	Port: jsii.String("port"),
//   	Protocol: jsii.String("protocol"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html
//
type CfnResolverRule_TargetAddressProperty struct {
	// One IPv4 address that you want to forward DNS queries to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-ip
	//
	Ip *string `field:"optional" json:"ip" yaml:"ip"`
	// One IPv6 address that you want to forward DNS queries to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-ipv6
	//
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// The port at `Ip` that you want to forward DNS queries to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-port
	//
	Port *string `field:"optional" json:"port" yaml:"port"`
	// The protocols for the Resolver endpoints. DoH-FIPS is applicable for inbound endpoints only.
	//
	// For an inbound endpoint you can apply the protocols as follows:
	//
	// - Do53 and DoH in combination.
	// - Do53 and DoH-FIPS in combination.
	// - Do53 alone.
	// - DoH alone.
	// - DoH-FIPS alone.
	// - None, which is treated as Do53.
	//
	// For an outbound endpoint you can apply the protocols as follows:
	//
	// - Do53 and DoH in combination.
	// - Do53 alone.
	// - DoH alone.
	// - None, which is treated as Do53.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}


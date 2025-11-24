package mixinsawsroute53resolver


// In a [CreateResolverRule](https://docs.aws.amazon.com/Route53/latest/APIReference/API_route53resolver_CreateResolverRule.html) request, an array of the IPs that you want to forward DNS queries to.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   targetAddressProperty := &TargetAddressProperty{
//   	Ip: jsii.String("ip"),
//   	Ipv6: jsii.String("ipv6"),
//   	Port: jsii.String("port"),
//   	Protocol: jsii.String("protocol"),
//   	ServerNameIndication: jsii.String("serverNameIndication"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html
//
type CfnResolverRulePropsMixin_TargetAddressProperty struct {
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
	// The protocols for the target address.
	//
	// The protocol you choose needs to be supported by the outbound endpoint of the Resolver rule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-protocol
	//
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// The Server Name Indication of the DoH server that you want to forward queries to.
	//
	// This is only used if the Protocol of the `TargetAddress` is `DoH` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-servernameindication
	//
	ServerNameIndication *string `field:"optional" json:"serverNameIndication" yaml:"serverNameIndication"`
}


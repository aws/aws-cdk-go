package previewawsec2mixins


// Describes the DNS options for an endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dnsOptionsSpecificationProperty := &DnsOptionsSpecificationProperty{
//   	DnsRecordIpType: jsii.String("dnsRecordIpType"),
//   	PrivateDnsOnlyForInboundResolverEndpoint: jsii.String("privateDnsOnlyForInboundResolverEndpoint"),
//   	PrivateDnsPreference: jsii.String("privateDnsPreference"),
//   	PrivateDnsSpecifiedDomains: []*string{
//   		jsii.String("privateDnsSpecifiedDomains"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html
//
type CfnVPCEndpointPropsMixin_DnsOptionsSpecificationProperty struct {
	// The DNS records created for the endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html#cfn-ec2-vpcendpoint-dnsoptionsspecification-dnsrecordiptype
	//
	DnsRecordIpType *string `field:"optional" json:"dnsRecordIpType" yaml:"dnsRecordIpType"`
	// Indicates whether to enable private DNS only for inbound endpoints.
	//
	// This option is available only for services that support both gateway and interface endpoints. It routes traffic that originates from the VPC to the gateway endpoint and traffic that originates from on-premises to the interface endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html#cfn-ec2-vpcendpoint-dnsoptionsspecification-privatednsonlyforinboundresolverendpoint
	//
	PrivateDnsOnlyForInboundResolverEndpoint *string `field:"optional" json:"privateDnsOnlyForInboundResolverEndpoint" yaml:"privateDnsOnlyForInboundResolverEndpoint"`
	// The preference for which private domains have a private hosted zone created for and associated with the specified VPC.
	//
	// Only supported when private DNS is enabled and when the VPC endpoint type is ServiceNetwork or Resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html#cfn-ec2-vpcendpoint-dnsoptionsspecification-privatednspreference
	//
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Indicates which of the private domains to create private hosted zones for and associate with the specified VPC.
	//
	// Only supported when private DNS is enabled and the private DNS preference is `VERIFIED_DOMAINS_AND_SPECIFIED_DOMAINS` or `SPECIFIED_DOMAINS_ONLY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-vpcendpoint-dnsoptionsspecification.html#cfn-ec2-vpcendpoint-dnsoptionsspecification-privatednsspecifieddomains
	//
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}


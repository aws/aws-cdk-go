package previewawsvpclatticemixins


// The DNS configuration options.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dnsOptionsProperty := &DnsOptionsProperty{
//   	PrivateDnsPreference: jsii.String("privateDnsPreference"),
//   	PrivateDnsSpecifiedDomains: []*string{
//   		jsii.String("privateDnsSpecifiedDomains"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetworkvpcassociation-dnsoptions.html
//
type CfnServiceNetworkVpcAssociationPropsMixin_DnsOptionsProperty struct {
	// The preference for which private domains have a private hosted zone created for and associated with the specified VPC.
	//
	// Only supported when private DNS is enabled and when the VPC endpoint type is ServiceNetwork or Resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetworkvpcassociation-dnsoptions.html#cfn-vpclattice-servicenetworkvpcassociation-dnsoptions-privatednspreference
	//
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// Indicates which of the private domains to create private hosted zones for and associate with the specified VPC.
	//
	// Only supported when private DNS is enabled and the private DNS preference is `VERIFIED_DOMAINS_AND_SPECIFIED_DOMAINS` or `SPECIFIED_DOMAINS_ONLY` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetworkvpcassociation-dnsoptions.html#cfn-vpclattice-servicenetworkvpcassociation-dnsoptions-privatednsspecifieddomains
	//
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}


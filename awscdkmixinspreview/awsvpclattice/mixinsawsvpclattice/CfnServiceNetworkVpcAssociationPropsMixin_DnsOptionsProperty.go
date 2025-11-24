package mixinsawsvpclattice


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
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetworkvpcassociation-dnsoptions.html#cfn-vpclattice-servicenetworkvpcassociation-dnsoptions-privatednspreference
	//
	PrivateDnsPreference *string `field:"optional" json:"privateDnsPreference" yaml:"privateDnsPreference"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-vpclattice-servicenetworkvpcassociation-dnsoptions.html#cfn-vpclattice-servicenetworkvpcassociation-dnsoptions-privatednsspecifieddomains
	//
	PrivateDnsSpecifiedDomains *[]*string `field:"optional" json:"privateDnsSpecifiedDomains" yaml:"privateDnsSpecifiedDomains"`
}


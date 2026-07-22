package awscertificatemanager


// DNS-based prevalidation options for the domain validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   dnsPrevalidationOptionsProperty := &DnsPrevalidationOptionsProperty{
//   	DomainScope: &DomainScopeProperty{
//   		ExactDomain: jsii.String("exactDomain"),
//   		Subdomains: jsii.String("subdomains"),
//   		Wildcards: jsii.String("wildcards"),
//   	},
//   	HostedZoneId: jsii.String("hostedZoneId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-dnsprevalidationoptions.html
//
type CfnAcmeDomainValidationPropsMixin_DnsPrevalidationOptionsProperty struct {
	// Controls which certificate types are authorized to be issued for the domain via the ACME endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-dnsprevalidationoptions.html#cfn-certificatemanager-acmedomainvalidation-dnsprevalidationoptions-domainscope
	//
	DomainScope interface{} `field:"optional" json:"domainScope" yaml:"domainScope"`
	// The Route 53 hosted zone ID for automatic DNS record management.
	//
	// When provided, the service creates the validation DNS record on the customer's behalf.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-dnsprevalidationoptions.html#cfn-certificatemanager-acmedomainvalidation-dnsprevalidationoptions-hostedzoneid
	//
	HostedZoneId *string `field:"optional" json:"hostedZoneId" yaml:"hostedZoneId"`
}


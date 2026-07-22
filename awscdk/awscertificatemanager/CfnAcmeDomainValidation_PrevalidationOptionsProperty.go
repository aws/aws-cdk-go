package awscertificatemanager


// Prevalidation method configuration.
//
// Currently only DNS-based prevalidation is supported.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   prevalidationOptionsProperty := &PrevalidationOptionsProperty{
//   	DnsPrevalidation: &DnsPrevalidationOptionsProperty{
//   		DomainScope: &DomainScopeProperty{
//   			ExactDomain: jsii.String("exactDomain"),
//   			Subdomains: jsii.String("subdomains"),
//   			Wildcards: jsii.String("wildcards"),
//   		},
//   		HostedZoneId: jsii.String("hostedZoneId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-prevalidationoptions.html
//
type CfnAcmeDomainValidation_PrevalidationOptionsProperty struct {
	// DNS-based prevalidation options for the domain validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-prevalidationoptions.html#cfn-certificatemanager-acmedomainvalidation-prevalidationoptions-dnsprevalidation
	//
	DnsPrevalidation interface{} `field:"required" json:"dnsPrevalidation" yaml:"dnsPrevalidation"`
}


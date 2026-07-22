package awscertificatemanager


// Controls which certificate types are authorized to be issued for the domain via the ACME endpoint.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   domainScopeProperty := &DomainScopeProperty{
//   	ExactDomain: jsii.String("exactDomain"),
//   	Subdomains: jsii.String("subdomains"),
//   	Wildcards: jsii.String("wildcards"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-domainscope.html
//
type CfnAcmeDomainValidation_DomainScopeProperty struct {
	// Whether certificates may be issued for the exact domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-domainscope.html#cfn-certificatemanager-acmedomainvalidation-domainscope-exactdomain
	//
	ExactDomain *string `field:"optional" json:"exactDomain" yaml:"exactDomain"`
	// Whether certificates may be issued for subdomains of the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-domainscope.html#cfn-certificatemanager-acmedomainvalidation-domainscope-subdomains
	//
	Subdomains *string `field:"optional" json:"subdomains" yaml:"subdomains"`
	// Whether wildcard certificates may be issued for the domain.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-acmedomainvalidation-domainscope.html#cfn-certificatemanager-acmedomainvalidation-domainscope-wildcards
	//
	Wildcards *string `field:"optional" json:"wildcards" yaml:"wildcards"`
}


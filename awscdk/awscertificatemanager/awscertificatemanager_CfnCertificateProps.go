package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk"
)

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &cfnCertificateProps{
//   	domainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	certificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	certificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   	domainValidationOptions: []interface{}{
//   		&domainValidationOptionProperty{
//   			domainName: jsii.String("domainName"),
//
//   			// the properties below are optional
//   			hostedZoneId: jsii.String("hostedZoneId"),
//   			validationDomain: jsii.String("validationDomain"),
//   		},
//   	},
//   	subjectAlternativeNames: []*string{
//   		jsii.String("subjectAlternativeNames"),
//   	},
//   	tags: []cfnTag{
//   		&cfnTag{
//   			key: jsii.String("key"),
//   			value: jsii.String("value"),
//   		},
//   	},
//   	validationMethod: jsii.String("validationMethod"),
//   }
//
type CfnCertificateProps struct {
	// The fully qualified domain name (FQDN), such as www.example.com, with which you want to secure an ACM certificate. Use an asterisk (*) to create a wildcard certificate that protects several sites in the same domain. For example, `*.example.com` protects `www.example.com` , `site.example.com` , and `images.example.com.`.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.
	//
	// If you do not provide an ARN and you are trying to request a private certificate, ACM will attempt to issue a public certificate. For more information about private CAs, see the [AWS Certificate Manager Private Certificate Authority (PCA)](https://docs.aws.amazon.com/acm-pca/latest/userguide/PcaWelcome.html) user guide. The ARN must have the following form:
	//
	// `arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012`.
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// You can opt out of certificate transparency logging by specifying the `DISABLED` option. Opt in by specifying `ENABLED` .
	//
	// If you do not specify a certificate transparency logging preference on a new CloudFormation template, or if you remove the logging preference from an existing template, this is the same as explicitly enabling the preference.
	//
	// Changing the certificate transparency logging preference will update the existing resource by calling `UpdateCertificateOptions` on the certificate. This action will not create a new resource.
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Domain information that domain name registrars use to verify your identity.
	//
	// > In order for a AWS::CertificateManager::Certificate to be provisioned and validated in CloudFormation automatically, the `DomainName` property needs to be identical to one of the `DomainName` property supplied in DomainValidationOptions, if the ValidationMethod is **DNS**. Failing to keep them like-for-like will result in failure to create the domain validation records in Route53.
	DomainValidationOptions interface{} `field:"optional" json:"domainValidationOptions" yaml:"domainValidationOptions"`
	// Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate.
	//
	// For example, you can add www.example.net to a certificate for which the `DomainName` field is www.example.com if users can reach your site by using either name.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Key-value pairs that can identify the certificate.
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The method you want to use to validate that you own or control the domain associated with a public certificate.
	//
	// You can [validate with DNS](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html) or [validate with email](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html) . We recommend that you use DNS validation.
	//
	// If not specified, this property defaults to email validation.
	ValidationMethod *string `field:"optional" json:"validationMethod" yaml:"validationMethod"`
}


package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCertificate`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnCertificateProps := &CfnCertificateProps{
//   	DomainName: jsii.String("domainName"),
//
//   	// the properties below are optional
//   	CertificateAuthorityArn: jsii.String("certificateAuthorityArn"),
//   	CertificateTransparencyLoggingPreference: jsii.String("certificateTransparencyLoggingPreference"),
//   	DomainValidationOptions: []interface{}{
//   		&DomainValidationOptionProperty{
//   			DomainName: jsii.String("domainName"),
//
//   			// the properties below are optional
//   			HostedZoneId: jsii.String("hostedZoneId"),
//   			ValidationDomain: jsii.String("validationDomain"),
//   		},
//   	},
//   	KeyAlgorithm: jsii.String("keyAlgorithm"),
//   	SubjectAlternativeNames: []*string{
//   		jsii.String("subjectAlternativeNames"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   	ValidationMethod: jsii.String("validationMethod"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html
//
type CfnCertificateProps struct {
	// The fully qualified domain name (FQDN), such as www.example.com, with which you want to secure an ACM certificate. Use an asterisk (*) to create a wildcard certificate that protects several sites in the same domain. For example, `*.example.com` protects `www.example.com` , `site.example.com` , and `images.example.com.`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainname
	//
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Amazon Resource Name (ARN) of the private certificate authority (CA) that will be used to issue the certificate.
	//
	// If you do not provide an ARN and you are trying to request a private certificate, ACM will attempt to issue a public certificate. For more information about private CAs, see the [AWS Private Certificate Authority](https://docs.aws.amazon.com/privateca/latest/userguide/PcaWelcome.html) user guide. The ARN must have the following form:
	//
	// `arn:aws:acm-pca:region:account:certificate-authority/12345678-1234-1234-1234-123456789012`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-certificateauthorityarn
	//
	CertificateAuthorityArn *string `field:"optional" json:"certificateAuthorityArn" yaml:"certificateAuthorityArn"`
	// You can opt out of certificate transparency logging by specifying the `DISABLED` option. Opt in by specifying `ENABLED` .
	//
	// If you do not specify a certificate transparency logging preference on a new CloudFormation template, or if you remove the logging preference from an existing template, this is the same as explicitly enabling the preference.
	//
	// Changing the certificate transparency logging preference will update the existing resource by calling `UpdateCertificateOptions` on the certificate. This action will not create a new resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-certificatetransparencyloggingpreference
	//
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Domain information that domain name registrars use to verify your identity.
	//
	// > In order for a AWS::CertificateManager::Certificate to be provisioned and validated in CloudFormation automatically, the `DomainName` property needs to be identical to one of the `DomainName` property supplied in DomainValidationOptions, if the ValidationMethod is **DNS**. Failing to keep them like-for-like will result in failure to create the domain validation records in Route53.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainvalidationoptions
	//
	DomainValidationOptions interface{} `field:"optional" json:"domainValidationOptions" yaml:"domainValidationOptions"`
	// Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data.
	//
	// RSA is the default key algorithm for ACM certificates. Elliptic Curve Digital Signature Algorithm (ECDSA) keys are smaller, offering security comparable to RSA keys but with greater computing efficiency. However, ECDSA is not supported by all network clients. Some AWS services may require RSA keys, or only support ECDSA keys of a particular size, while others allow the use of either RSA and ECDSA keys to ensure that compatibility is not broken. Check the requirements for the AWS service where you plan to deploy your certificate. For more information about selecting an algorithm, see [Key algorithms](https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms) .
	//
	// > Algorithms supported for an ACM certificate request include:
	// >
	// > - `RSA_2048`
	// > - `EC_prime256v1`
	// > - `EC_secp384r1`
	// >
	// > Other listed algorithms are for imported certificates only. > When you request a private PKI certificate signed by a CA from AWS Private CA, the specified signing algorithm family (RSA or ECDSA) must match the algorithm family of the CA's secret key.
	//
	// Default: RSA_2048.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-keyalgorithm
	//
	KeyAlgorithm *string `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Additional FQDNs to be included in the Subject Alternative Name extension of the ACM certificate.
	//
	// For example, you can add www.example.net to a certificate for which the `DomainName` field is www.example.com if users can reach your site by using either name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-subjectalternativenames
	//
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Key-value pairs that can identify the certificate.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
	// The method you want to use to validate that you own or control the domain associated with a public certificate.
	//
	// You can [validate with DNS](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-dns.html) or [validate with email](https://docs.aws.amazon.com/acm/latest/userguide/gs-acm-validate-email.html) . We recommend that you use DNS validation.
	//
	// If not specified, this property defaults to email validation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-validationmethod
	//
	ValidationMethod *string `field:"optional" json:"validationMethod" yaml:"validationMethod"`
}


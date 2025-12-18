package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsacmpca"
)

// Properties for your private certificate.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//
//   acm.NewPrivateCertificate(this, jsii.String("PrivateCertificate"), &PrivateCertificateProps{
//   	DomainName: jsii.String("test.example.com"),
//   	SubjectAlternativeNames: []*string{
//   		jsii.String("cool.example.com"),
//   		jsii.String("test.example.net"),
//   	},
//   	 // optional
//   	CertificateAuthority: acmpca.CertificateAuthority_FromCertificateAuthorityArn(this, jsii.String("CA"), jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/023077d8-2bfa-4eb0-8f22-05c96deade77")),
//   	KeyAlgorithm: acm.KeyAlgorithm_RSA_2048(),
//   })
//
type PrivateCertificateProps struct {
	// Private certificate authority (CA) that will be used to issue the certificate.
	CertificateAuthority interfacesawsacmpca.ICertificateAuthorityRef `field:"required" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Fully-qualified domain name to request a private certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Enable or disable export of this certificate.
	//
	// If you issue an exportable public certificate, there is a charge at certificate issuance and again when the certificate renews.
	// Ref: https://aws.amazon.com/certificate-manager/pricing
	// Default: false.
	//
	AllowExport *bool `field:"optional" json:"allowExport" yaml:"allowExport"`
	// Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data.
	//
	// When you request a private PKI certificate signed by a CA from AWS Private CA, the specified signing algorithm family
	// (RSA or ECDSA) must match the algorithm family of the CA's secret key.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms.title
	//
	// Default: KeyAlgorithm.RSA_2048
	//
	KeyAlgorithm KeyAlgorithm `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Alternative domain names on your private certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	// Default: - No additional FQDNs will be included as alternative domain names.
	//
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}


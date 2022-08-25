package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca"
)

// Properties for your private certificate.
//
// Example:
//   import acmpca "github.com/aws/aws-cdk-go/awscdk"
//
//
//   acm.NewPrivateCertificate(this, jsii.String("PrivateCertificate"), &privateCertificateProps{
//   	domainName: jsii.String("test.example.com"),
//   	subjectAlternativeNames: []*string{
//   		jsii.String("cool.example.com"),
//   		jsii.String("test.example.net"),
//   	},
//   	 // optional
//   	certificateAuthority: acmpca.certificateAuthority.fromCertificateAuthorityArn(this, jsii.String("CA"), jsii.String("arn:aws:acm-pca:us-east-1:123456789012:certificate-authority/023077d8-2bfa-4eb0-8f22-05c96deade77")),
//   })
//
type PrivateCertificateProps struct {
	// Private certificate authority (CA) that will be used to issue the certificate.
	CertificateAuthority awsacmpca.ICertificateAuthority `field:"required" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Fully-qualified domain name to request a private certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Alternative domain names on your private certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
}


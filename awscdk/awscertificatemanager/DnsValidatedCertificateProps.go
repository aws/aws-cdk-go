package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties to create a DNS validated certificate managed by AWS Certificate Manager.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var certificateValidation certificateValidation
//   var hostedZone hostedZone
//   var keyAlgorithm keyAlgorithm
//   var role role
//
//   dnsValidatedCertificateProps := &DnsValidatedCertificateProps{
//   	DomainName: jsii.String("domainName"),
//   	HostedZone: hostedZone,
//
//   	// the properties below are optional
//   	CertificateName: jsii.String("certificateName"),
//   	CleanupRoute53Records: jsii.Boolean(false),
//   	CustomResourceRole: role,
//   	KeyAlgorithm: keyAlgorithm,
//   	Region: jsii.String("region"),
//   	Route53Endpoint: jsii.String("route53Endpoint"),
//   	SubjectAlternativeNames: []*string{
//   		jsii.String("subjectAlternativeNames"),
//   	},
//   	TransparencyLoggingEnabled: jsii.Boolean(false),
//   	Validation: certificateValidation,
//   }
//
type DnsValidatedCertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Certificate name.
	//
	// Since the Certificate resource doesn't support providing a physical name, the value provided here will be recorded in the `Name` tag.
	// Default: the full, absolute path of this construct.
	//
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Specifies the algorithm of the public and private key pair that your certificate uses to encrypt data.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/acm-certificate.html#algorithms.title
	//
	// Default: KeyAlgorithm.RSA_2048
	//
	KeyAlgorithm KeyAlgorithm `field:"optional" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	// Default: - No additional FQDNs will be included as alternative domain names.
	//
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// Enable or disable transparency logging for this certificate.
	//
	// Once a certificate has been logged, it cannot be removed from the log.
	// Opting out at that point will have no effect. If you opt out of logging
	// when you request a certificate and then choose later to opt back in,
	// your certificate will not be logged until it is renewed.
	// If you want the certificate to be logged immediately, we recommend that you issue a new one.
	// See: https://docs.aws.amazon.com/acm/latest/userguide/acm-bestpractices.html#best-practices-transparency
	//
	// Default: true.
	//
	TransparencyLoggingEnabled *bool `field:"optional" json:"transparencyLoggingEnabled" yaml:"transparencyLoggingEnabled"`
	// How to validate this certificate.
	// Default: CertificateValidation.fromEmail()
	//
	Validation CertificateValidation `field:"optional" json:"validation" yaml:"validation"`
	// Route 53 Hosted Zone used to perform DNS validation of the request.
	//
	// The zone
	// must be authoritative for the domain name specified in the Certificate Request.
	HostedZone awsroute53.IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// When set to true, when the DnsValidatedCertificate is deleted, the associated Route53 validation records are removed.
	//
	// CAUTION: If multiple certificates share the same domains (and same validation records),
	// this can cause the other certificates to fail renewal and/or not validate.
	// Not recommended for production use.
	// Default: false.
	//
	CleanupRoute53Records *bool `field:"optional" json:"cleanupRoute53Records" yaml:"cleanupRoute53Records"`
	// Role to use for the custom resource that creates the validated certificate.
	// Default: - A new role will be created.
	//
	CustomResourceRole awsiam.IRole `field:"optional" json:"customResourceRole" yaml:"customResourceRole"`
	// AWS region that will host the certificate.
	//
	// This is needed especially
	// for certificates used for CloudFront distributions, which require the region
	// to be us-east-1.
	// Default: the region the stack is deployed in.
	//
	Region *string `field:"optional" json:"region" yaml:"region"`
	// An endpoint of Route53 service, which is not necessary as AWS SDK could figure out the right endpoints for most regions, but for some regions such as those in aws-cn partition, the default endpoint is not working now, hence the right endpoint need to be specified through this prop.
	//
	// Route53 is not been officially launched in China, it is only available for AWS
	// internal accounts now. To make DnsValidatedCertificate work for internal accounts
	// now, a special endpoint needs to be provided.
	// Default: - The AWS SDK will determine the Route53 endpoint to use based on region.
	//
	Route53Endpoint *string `field:"optional" json:"route53Endpoint" yaml:"route53Endpoint"`
}


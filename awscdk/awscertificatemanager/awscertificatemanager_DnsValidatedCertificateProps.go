package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties to create a DNS validated certificate managed by AWS Certificate Manager.
//
// Example:
//   var myHostedZone hostedZone
//
//   acm.NewDnsValidatedCertificate(this, jsii.String("CrossRegionCertificate"), &dnsValidatedCertificateProps{
//   	domainName: jsii.String("hello.example.com"),
//   	hostedZone: myHostedZone,
//   	region: jsii.String("us-east-1"),
//   })
//
type DnsValidatedCertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// The Certifcate name.
	//
	// Since the Certifcate resource doesn't support providing a physical name, the value provided here will be recorded in the `Name` tag.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
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
	TransparencyLoggingEnabled *bool `field:"optional" json:"transparencyLoggingEnabled" yaml:"transparencyLoggingEnabled"`
	// How to validate this certificate.
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
	CleanupRoute53Records *bool `field:"optional" json:"cleanupRoute53Records" yaml:"cleanupRoute53Records"`
	// Role to use for the custom resource that creates the validated certificate.
	CustomResourceRole awsiam.IRole `field:"optional" json:"customResourceRole" yaml:"customResourceRole"`
	// AWS region that will host the certificate.
	//
	// This is needed especially
	// for certificates used for CloudFront distributions, which require the region
	// to be us-east-1.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// An endpoint of Route53 service, which is not necessary as AWS SDK could figure out the right endpoints for most regions, but for some regions such as those in aws-cn partition, the default endpoint is not working now, hence the right endpoint need to be specified through this prop.
	//
	// Route53 is not been officially launched in China, it is only available for AWS
	// internal accounts now. To make DnsValidatedCertificate work for internal accounts
	// now, a special endpoint needs to be provided.
	Route53Endpoint *string `field:"optional" json:"route53Endpoint" yaml:"route53Endpoint"`
}


package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
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
// Experimental.
type DnsValidatedCertificateProps struct {
	// Fully-qualified domain name to request a certificate for.
	//
	// May contain wildcards, such as ``*.domain.com``.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Alternative domain names on your certificate.
	//
	// Use this to register alternative domain names that represent the same site.
	// Experimental.
	SubjectAlternativeNames *[]*string `field:"optional" json:"subjectAlternativeNames" yaml:"subjectAlternativeNames"`
	// How to validate this certificate.
	// Experimental.
	Validation CertificateValidation `field:"optional" json:"validation" yaml:"validation"`
	// What validation domain to use for every requested domain.
	//
	// Has to be a superdomain of the requested domain.
	// Deprecated: use `validation` instead.
	ValidationDomains *map[string]*string `field:"optional" json:"validationDomains" yaml:"validationDomains"`
	// Validation method used to assert domain ownership.
	// Deprecated: use `validation` instead.
	ValidationMethod ValidationMethod `field:"optional" json:"validationMethod" yaml:"validationMethod"`
	// Route 53 Hosted Zone used to perform DNS validation of the request.
	//
	// The zone
	// must be authoritative for the domain name specified in the Certificate Request.
	// Experimental.
	HostedZone awsroute53.IHostedZone `field:"required" json:"hostedZone" yaml:"hostedZone"`
	// When set to true, when the DnsValidatedCertificate is deleted, the associated Route53 validation records are removed.
	//
	// CAUTION: If multiple certificates share the same domains (and same validation records),
	// this can cause the other certificates to fail renewal and/or not validate.
	// Not recommended for production use.
	// Experimental.
	CleanupRoute53Records *bool `field:"optional" json:"cleanupRoute53Records" yaml:"cleanupRoute53Records"`
	// Role to use for the custom resource that creates the validated certificate.
	// Experimental.
	CustomResourceRole awsiam.IRole `field:"optional" json:"customResourceRole" yaml:"customResourceRole"`
	// AWS region that will host the certificate.
	//
	// This is needed especially
	// for certificates used for CloudFront distributions, which require the region
	// to be us-east-1.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// An endpoint of Route53 service, which is not necessary as AWS SDK could figure out the right endpoints for most regions, but for some regions such as those in aws-cn partition, the default endpoint is not working now, hence the right endpoint need to be specified through this prop.
	//
	// Route53 is not been officially launched in China, it is only available for AWS
	// internal accounts now. To make DnsValidatedCertificate work for internal accounts
	// now, a special endpoint needs to be provided.
	// Experimental.
	Route53Endpoint *string `field:"optional" json:"route53Endpoint" yaml:"route53Endpoint"`
}


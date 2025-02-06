package awscertificatemanager


// Properties for your certificate.
//
// Example:
//   // To use your own domain name in a Distribution, you must associate a certificate
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import route53 "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   var myBucket bucket
//
//   myCertificate := acm.NewCertificate(this, jsii.String("mySiteCert"), &CertificateProps{
//   	DomainName: jsii.String("www.example.com"),
//   	Validation: acm.CertificateValidation_FromDns(hostedZone),
//   })
//   cloudfront.NewDistribution(this, jsii.String("myDist"), &DistributionProps{
//   	DefaultBehavior: &BehaviorOptions{
//   		Origin: origins.NewS3Origin(myBucket),
//   	},
//   	DomainNames: []*string{
//   		jsii.String("www.example.com"),
//   	},
//   	Certificate: myCertificate,
//   })
//
type CertificateProps struct {
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
}


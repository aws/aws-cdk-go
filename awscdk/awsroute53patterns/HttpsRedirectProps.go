package awsroute53patterns

import (
	"github.com/aws/aws-cdk-go/awscdk/awscertificatemanager"
	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

// Properties to configure an HTTPS Redirect.
//
// Example:
//   patterns.NewHttpsRedirect(this, jsii.String("Redirect"), &HttpsRedirectProps{
//   	RecordNames: []*string{
//   		jsii.String("foo.example.com"),
//   	},
//   	TargetDomain: jsii.String("bar.example.com"),
//   	Zone: route53.HostedZone_FromHostedZoneAttributes(this, jsii.String("HostedZone"), &HostedZoneAttributes{
//   		HostedZoneId: jsii.String("ID"),
//   		ZoneName: jsii.String("example.com"),
//   	}),
//   })
//
// Experimental.
type HttpsRedirectProps struct {
	// The redirect target fully qualified domain name (FQDN).
	//
	// An alias record
	// will be created that points to your CloudFront distribution. Root domain
	// or sub-domain can be supplied.
	// Experimental.
	TargetDomain *string `field:"required" json:"targetDomain" yaml:"targetDomain"`
	// Hosted zone of the domain which will be used to create alias record(s) from domain names in the hosted zone to the target domain.
	//
	// The hosted zone must
	// contain entries for the domain name(s) supplied through `recordNames` that
	// will redirect to the target domain.
	//
	// Domain names in the hosted zone can include a specific domain (example.com)
	// and its subdomains (acme.example.com, zenith.example.com).
	// Experimental.
	Zone awsroute53.IHostedZone `field:"required" json:"zone" yaml:"zone"`
	// The AWS Certificate Manager (ACM) certificate that will be associated with the CloudFront distribution that will be created.
	//
	// If provided, the certificate must be
	// stored in us-east-1 (N. Virginia)
	// Experimental.
	Certificate awscertificatemanager.ICertificate `field:"optional" json:"certificate" yaml:"certificate"`
	// The domain names that will redirect to `targetDomain`.
	// Experimental.
	RecordNames *[]*string `field:"optional" json:"recordNames" yaml:"recordNames"`
}


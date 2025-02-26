package awscertificatemanager

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// Properties for certificate validation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var hostedZone hostedZone
//
//   certificationValidationProps := &CertificationValidationProps{
//   	HostedZone: hostedZone,
//   	HostedZones: map[string]iHostedZone{
//   		"hostedZonesKey": hostedZone,
//   	},
//   	Method: awscdk.Aws_certificatemanager.ValidationMethod_EMAIL,
//   	ValidationDomains: map[string]*string{
//   		"validationDomainsKey": jsii.String("validationDomains"),
//   	},
//   }
//
type CertificationValidationProps struct {
	// Hosted zone to use for DNS validation.
	// Default: - use email validation.
	//
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
	// A map of hosted zones to use for DNS validation.
	// Default: - use `hostedZone`.
	//
	HostedZones *map[string]awsroute53.IHostedZone `field:"optional" json:"hostedZones" yaml:"hostedZones"`
	// Validation method.
	// Default: ValidationMethod.EMAIL
	//
	Method ValidationMethod `field:"optional" json:"method" yaml:"method"`
	// Validation domains to use for email validation.
	// Default: - Apex domain.
	//
	ValidationDomains *map[string]*string `field:"optional" json:"validationDomains" yaml:"validationDomains"`
}


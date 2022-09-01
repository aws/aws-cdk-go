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
//   certificationValidationProps := &certificationValidationProps{
//   	hostedZone: hostedZone,
//   	hostedZones: map[string]iHostedZone{
//   		"hostedZonesKey": hostedZone,
//   	},
//   	method: awscdk.Aws_certificatemanager.validationMethod_EMAIL,
//   	validationDomains: map[string]*string{
//   		"validationDomainsKey": jsii.String("validationDomains"),
//   	},
//   }
//
type CertificationValidationProps struct {
	// Hosted zone to use for DNS validation.
	HostedZone awsroute53.IHostedZone `field:"optional" json:"hostedZone" yaml:"hostedZone"`
	// A map of hosted zones to use for DNS validation.
	HostedZones *map[string]awsroute53.IHostedZone `field:"optional" json:"hostedZones" yaml:"hostedZones"`
	// Validation method.
	Method ValidationMethod `field:"optional" json:"method" yaml:"method"`
	// Validation domains to use for email validation.
	ValidationDomains *map[string]*string `field:"optional" json:"validationDomains" yaml:"validationDomains"`
}


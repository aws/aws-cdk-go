package awscertificatemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsroute53"
)

// How to validate a certificate.
//
// Example:
//   exampleCom := route53.NewHostedZone(this, jsii.String("ExampleCom"), &HostedZoneProps{
//   	ZoneName: jsii.String("example.com"),
//   })
//   exampleNet := route53.NewHostedZone(this, jsii.String("ExampleNet"), &HostedZoneProps{
//   	ZoneName: jsii.String("example.net"),
//   })
//
//   cert := acm.NewCertificate(this, jsii.String("Certificate"), &CertificateProps{
//   	DomainName: jsii.String("test.example.com"),
//   	SubjectAlternativeNames: []*string{
//   		jsii.String("cool.example.com"),
//   		jsii.String("test.example.net"),
//   	},
//   	Validation: acm.CertificateValidation_FromDnsMultiZone(map[string]iHostedZone{
//   		"test.example.com": exampleCom,
//   		"cool.example.com": exampleCom,
//   		"test.example.net": exampleNet,
//   	}),
//   })
//
type CertificateValidation interface {
	// The validation method.
	Method() ValidationMethod
	// Certification validation properties.
	Props() *CertificationValidationProps
}

// The jsii proxy struct for CertificateValidation
type jsiiProxy_CertificateValidation struct {
	_ byte // padding
}

func (j *jsiiProxy_CertificateValidation) Method() ValidationMethod {
	var returns ValidationMethod
	_jsii_.Get(
		j,
		"method",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CertificateValidation) Props() *CertificationValidationProps {
	var returns *CertificationValidationProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}


// Validate the certificate with DNS.
//
// IMPORTANT: If `hostedZone` is not specified, DNS records must be added
// manually and the stack will not complete creating until the records are
// added.
func CertificateValidation_FromDns(hostedZone awsroute53.IHostedZone) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CertificateValidation",
		"fromDns",
		[]interface{}{hostedZone},
		&returns,
	)

	return returns
}

// Validate the certificate with automatically created DNS records in multiple Amazon Route 53 hosted zones.
func CertificateValidation_FromDnsMultiZone(hostedZones *map[string]awsroute53.IHostedZone) CertificateValidation {
	_init_.Initialize()

	if err := validateCertificateValidation_FromDnsMultiZoneParameters(hostedZones); err != nil {
		panic(err)
	}
	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CertificateValidation",
		"fromDnsMultiZone",
		[]interface{}{hostedZones},
		&returns,
	)

	return returns
}

// Validate the certificate with Email.
//
// IMPORTANT: if you are creating a certificate as part of your stack, the stack
// will not complete creating until you read and follow the instructions in the
// email that you will receive.
//
// ACM will send validation emails to the following addresses:
//
//  admin@domain.com
//  administrator@domain.com
//  hostmaster@domain.com
//  postmaster@domain.com
//  webmaster@domain.com
//
// For every domain that you register.
func CertificateValidation_FromEmail(validationDomains *map[string]*string) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_certificatemanager.CertificateValidation",
		"fromEmail",
		[]interface{}{validationDomains},
		&returns,
	)

	return returns
}


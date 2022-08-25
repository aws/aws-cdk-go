package awscertificatemanager

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsroute53"
)

// How to validate a certificate.
//
// Example:
//   myHostedZone := route53.NewHostedZone(this, jsii.String("HostedZone"), &hostedZoneProps{
//   	zoneName: jsii.String("example.com"),
//   })
//   acm.NewCertificate(this, jsii.String("Certificate"), &certificateProps{
//   	domainName: jsii.String("hello.example.com"),
//   	validation: acm.certificateValidation.fromDns(myHostedZone),
//   })
//
// Experimental.
type CertificateValidation interface {
	// The validation method.
	// Experimental.
	Method() ValidationMethod
	// Certification validation properties.
	// Experimental.
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
// Experimental.
func CertificateValidation_FromDns(hostedZone awsroute53.IHostedZone) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"monocdk.aws_certificatemanager.CertificateValidation",
		"fromDns",
		[]interface{}{hostedZone},
		&returns,
	)

	return returns
}

// Validate the certificate with automatically created DNS records in multiple Amazon Route 53 hosted zones.
// Experimental.
func CertificateValidation_FromDnsMultiZone(hostedZones *map[string]awsroute53.IHostedZone) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"monocdk.aws_certificatemanager.CertificateValidation",
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
//   admin@domain.com
//   administrator@domain.com
//   hostmaster@domain.com
//   postmaster@domain.com
//   webmaster@domain.com
//
// For every domain that you register.
// Experimental.
func CertificateValidation_FromEmail(validationDomains *map[string]*string) CertificateValidation {
	_init_.Initialize()

	var returns CertificateValidation

	_jsii_.StaticInvoke(
		"monocdk.aws_certificatemanager.CertificateValidation",
		"fromEmail",
		[]interface{}{validationDomains},
		&returns,
	)

	return returns
}


package awselasticloadbalancingv2

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscertificatemanager"
)

// A certificate source for an ELBv2 listener.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   listenerCertificate := awscdk.Aws_elasticloadbalancingv2.ListenerCertificate_FromArn(jsii.String("certificateArn"))
//
type ListenerCertificate interface {
	IListenerCertificate
	// The ARN of the certificate to use.
	CertificateArn() *string
}

// The jsii proxy struct for ListenerCertificate
type jsiiProxy_ListenerCertificate struct {
	jsiiProxy_IListenerCertificate
}

func (j *jsiiProxy_ListenerCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}


func NewListenerCertificate(certificateArn *string) ListenerCertificate {
	_init_.Initialize()

	if err := validateNewListenerCertificateParameters(certificateArn); err != nil {
		panic(err)
	}
	j := jsiiProxy_ListenerCertificate{}

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		[]interface{}{certificateArn},
		&j,
	)

	return &j
}

func NewListenerCertificate_Override(l ListenerCertificate, certificateArn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		[]interface{}{certificateArn},
		l,
	)
}

// Use any certificate, identified by its ARN, as a listener certificate.
func ListenerCertificate_FromArn(certificateArn *string) ListenerCertificate {
	_init_.Initialize()

	if err := validateListenerCertificate_FromArnParameters(certificateArn); err != nil {
		panic(err)
	}
	var returns ListenerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		"fromArn",
		[]interface{}{certificateArn},
		&returns,
	)

	return returns
}

// Use an ACM certificate as a listener certificate.
func ListenerCertificate_FromCertificateManager(acmCertificate awscertificatemanager.ICertificate) ListenerCertificate {
	_init_.Initialize()

	if err := validateListenerCertificate_FromCertificateManagerParameters(acmCertificate); err != nil {
		panic(err)
	}
	var returns ListenerCertificate

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_elasticloadbalancingv2.ListenerCertificate",
		"fromCertificateManager",
		[]interface{}{acmCertificate},
		&returns,
	)

	return returns
}


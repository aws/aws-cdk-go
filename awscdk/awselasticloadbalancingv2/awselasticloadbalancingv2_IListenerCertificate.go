package awselasticloadbalancingv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// A certificate source for an ELBv2 listener.
type IListenerCertificate interface {
	// The ARN of the certificate to use.
	CertificateArn() *string
}

// The jsii proxy for IListenerCertificate
type jsiiProxy_IListenerCertificate struct {
	_ byte // padding
}

func (j *jsiiProxy_IListenerCertificate) CertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateArn",
		&returns,
	)
	return returns
}


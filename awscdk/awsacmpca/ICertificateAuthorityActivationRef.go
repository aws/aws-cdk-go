package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAuthorityActivation.
// Experimental.
type ICertificateAuthorityActivationRef interface {
	constructs.IConstruct
	// A reference to a CertificateAuthorityActivation resource.
	// Experimental.
	CertificateAuthorityActivationRef() *CertificateAuthorityActivationReference
}

// The jsii proxy for ICertificateAuthorityActivationRef
type jsiiProxy_ICertificateAuthorityActivationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICertificateAuthorityActivationRef) CertificateAuthorityActivationRef() *CertificateAuthorityActivationReference {
	var returns *CertificateAuthorityActivationReference
	_jsii_.Get(
		j,
		"certificateAuthorityActivationRef",
		&returns,
	)
	return returns
}


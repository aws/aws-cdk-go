package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CertificateAuthority.
// Experimental.
type ICertificateAuthorityRef interface {
	constructs.IConstruct
	// A reference to a CertificateAuthority resource.
	// Experimental.
	CertificateAuthorityRef() *CertificateAuthorityReference
}

// The jsii proxy for ICertificateAuthorityRef
type jsiiProxy_ICertificateAuthorityRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICertificateAuthorityRef) CertificateAuthorityRef() *CertificateAuthorityReference {
	var returns *CertificateAuthorityReference
	_jsii_.Get(
		j,
		"certificateAuthorityRef",
		&returns,
	)
	return returns
}


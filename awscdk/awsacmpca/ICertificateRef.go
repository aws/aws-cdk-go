package awsacmpca

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsacmpca/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Certificate.
// Experimental.
type ICertificateRef interface {
	constructs.IConstruct
	// A reference to a Certificate resource.
	// Experimental.
	CertificateRef() *CertificateReference
}

// The jsii proxy for ICertificateRef
type jsiiProxy_ICertificateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICertificateRef) CertificateRef() *CertificateReference {
	var returns *CertificateReference
	_jsii_.Get(
		j,
		"certificateRef",
		&returns,
	)
	return returns
}


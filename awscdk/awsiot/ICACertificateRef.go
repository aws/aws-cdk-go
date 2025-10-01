package awsiot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CACertificate.
// Experimental.
type ICACertificateRef interface {
	constructs.IConstruct
	// A reference to a CACertificate resource.
	// Experimental.
	CaCertificateRef() *CACertificateReference
}

// The jsii proxy for ICACertificateRef
type jsiiProxy_ICACertificateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICACertificateRef) CaCertificateRef() *CACertificateReference {
	var returns *CACertificateReference
	_jsii_.Get(
		j,
		"caCertificateRef",
		&returns,
	)
	return returns
}


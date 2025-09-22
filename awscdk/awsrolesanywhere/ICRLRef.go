package awsrolesanywhere

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsrolesanywhere/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a CRL.
// Experimental.
type ICRLRef interface {
	constructs.IConstruct
	// A reference to a CRL resource.
	// Experimental.
	CrlRef() *CRLReference
}

// The jsii proxy for ICRLRef
type jsiiProxy_ICRLRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ICRLRef) CrlRef() *CRLReference {
	var returns *CRLReference
	_jsii_.Get(
		j,
		"crlRef",
		&returns,
	)
	return returns
}


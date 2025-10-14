package awswaf

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswaf/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ByteMatchSet.
// Experimental.
type IByteMatchSetRef interface {
	constructs.IConstruct
	// A reference to a ByteMatchSet resource.
	// Experimental.
	ByteMatchSetRef() *ByteMatchSetReference
}

// The jsii proxy for IByteMatchSetRef
type jsiiProxy_IByteMatchSetRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IByteMatchSetRef) ByteMatchSetRef() *ByteMatchSetReference {
	var returns *ByteMatchSetReference
	_jsii_.Get(
		j,
		"byteMatchSetRef",
		&returns,
	)
	return returns
}


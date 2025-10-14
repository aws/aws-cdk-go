package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Multiplex.
// Experimental.
type IMultiplexRef interface {
	constructs.IConstruct
	// A reference to a Multiplex resource.
	// Experimental.
	MultiplexRef() *MultiplexReference
}

// The jsii proxy for IMultiplexRef
type jsiiProxy_IMultiplexRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMultiplexRef) MultiplexRef() *MultiplexReference {
	var returns *MultiplexReference
	_jsii_.Get(
		j,
		"multiplexRef",
		&returns,
	)
	return returns
}


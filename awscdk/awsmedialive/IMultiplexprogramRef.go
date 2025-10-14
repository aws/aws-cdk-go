package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Multiplexprogram.
// Experimental.
type IMultiplexprogramRef interface {
	constructs.IConstruct
	// A reference to a Multiplexprogram resource.
	// Experimental.
	MultiplexprogramRef() *MultiplexprogramReference
}

// The jsii proxy for IMultiplexprogramRef
type jsiiProxy_IMultiplexprogramRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMultiplexprogramRef) MultiplexprogramRef() *MultiplexprogramReference {
	var returns *MultiplexprogramReference
	_jsii_.Get(
		j,
		"multiplexprogramRef",
		&returns,
	)
	return returns
}


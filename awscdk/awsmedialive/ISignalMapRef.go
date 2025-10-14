package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalMap.
// Experimental.
type ISignalMapRef interface {
	constructs.IConstruct
	// A reference to a SignalMap resource.
	// Experimental.
	SignalMapRef() *SignalMapReference
}

// The jsii proxy for ISignalMapRef
type jsiiProxy_ISignalMapRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_ISignalMapRef) SignalMapRef() *SignalMapReference {
	var returns *SignalMapReference
	_jsii_.Get(
		j,
		"signalMapRef",
		&returns,
	)
	return returns
}


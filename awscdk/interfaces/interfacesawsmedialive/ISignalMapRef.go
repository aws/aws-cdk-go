package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SignalMap.
// Experimental.
type ISignalMapRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SignalMap resource.
	// Experimental.
	SignalMapRef() *SignalMapReference
}

// The jsii proxy for ISignalMapRef
type jsiiProxy_ISignalMapRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ISignalMapRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_ISignalMapRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISignalMapRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


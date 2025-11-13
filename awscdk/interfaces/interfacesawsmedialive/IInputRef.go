package interfacesawsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Input.
// Experimental.
type IInputRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Input resource.
	// Experimental.
	InputRef() *InputReference
}

// The jsii proxy for IInputRef
type jsiiProxy_IInputRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IInputRef) InputRef() *InputReference {
	var returns *InputReference
	_jsii_.Get(
		j,
		"inputRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInputRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


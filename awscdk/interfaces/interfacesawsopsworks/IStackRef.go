package interfacesawsopsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Stack.
// Experimental.
type IStackRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Stack resource.
	// Experimental.
	StackRef() *StackReference
}

// The jsii proxy for IStackRef
type jsiiProxy_IStackRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStackRef) StackRef() *StackReference {
	var returns *StackReference
	_jsii_.Get(
		j,
		"stackRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStackRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DaemonTaskDefinition.
// Experimental.
type IDaemonTaskDefinitionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a DaemonTaskDefinition resource.
	// Experimental.
	DaemonTaskDefinitionRef() *DaemonTaskDefinitionReference
}

// The jsii proxy for IDaemonTaskDefinitionRef
type jsiiProxy_IDaemonTaskDefinitionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IDaemonTaskDefinitionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IDaemonTaskDefinitionRef) DaemonTaskDefinitionRef() *DaemonTaskDefinitionReference {
	var returns *DaemonTaskDefinitionReference
	_jsii_.Get(
		j,
		"daemonTaskDefinitionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDaemonTaskDefinitionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IDaemonTaskDefinitionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


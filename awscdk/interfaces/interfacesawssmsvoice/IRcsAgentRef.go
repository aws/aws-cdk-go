package interfacesawssmsvoice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssmsvoice/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a RcsAgent.
// Experimental.
type IRcsAgentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a RcsAgent resource.
	// Experimental.
	RcsAgentRef() *RcsAgentReference
}

// The jsii proxy for IRcsAgentRef
type jsiiProxy_IRcsAgentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRcsAgentRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IRcsAgentRef) RcsAgentRef() *RcsAgentReference {
	var returns *RcsAgentReference
	_jsii_.Get(
		j,
		"rcsAgentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRcsAgentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRcsAgentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


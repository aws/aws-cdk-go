package interfacesawsdevopsagent

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdevopsagent/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentSpace.
// Experimental.
type IAgentSpaceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AgentSpace resource.
	// Experimental.
	AgentSpaceRef() *AgentSpaceReference
}

// The jsii proxy for IAgentSpaceRef
type jsiiProxy_IAgentSpaceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAgentSpaceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAgentSpaceRef) AgentSpaceRef() *AgentSpaceReference {
	var returns *AgentSpaceReference
	_jsii_.Get(
		j,
		"agentSpaceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentSpaceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentSpaceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


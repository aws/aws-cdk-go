package interfacesawsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentStatus.
// Experimental.
type IAgentStatusRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AgentStatus resource.
	// Experimental.
	AgentStatusRef() *AgentStatusReference
}

// The jsii proxy for IAgentStatusRef
type jsiiProxy_IAgentStatusRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAgentStatusRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_IAgentStatusRef) AgentStatusRef() *AgentStatusReference {
	var returns *AgentStatusReference
	_jsii_.Get(
		j,
		"agentStatusRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentStatusRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentStatusRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


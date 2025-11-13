package interfacesawsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentAlias.
// Experimental.
type IAgentAliasRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AgentAlias resource.
	// Experimental.
	AgentAliasRef() *AgentAliasReference
}

// The jsii proxy for IAgentAliasRef
type jsiiProxy_IAgentAliasRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAgentAliasRef) AgentAliasRef() *AgentAliasReference {
	var returns *AgentAliasReference
	_jsii_.Get(
		j,
		"agentAliasRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentAliasRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentAliasRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


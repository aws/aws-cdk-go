package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentAlias.
// Experimental.
type IAgentAliasRef interface {
	constructs.IConstruct
	// A reference to a AgentAlias resource.
	// Experimental.
	AgentAliasRef() *AgentAliasReference
}

// The jsii proxy for IAgentAliasRef
type jsiiProxy_IAgentAliasRef struct {
	internal.Type__constructsIConstruct
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


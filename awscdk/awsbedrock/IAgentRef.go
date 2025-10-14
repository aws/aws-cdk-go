package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Agent.
// Experimental.
type IAgentRef interface {
	constructs.IConstruct
	// A reference to a Agent resource.
	// Experimental.
	AgentRef() *AgentReference
}

// The jsii proxy for IAgentRef
type jsiiProxy_IAgentRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAgentRef) AgentRef() *AgentReference {
	var returns *AgentReference
	_jsii_.Get(
		j,
		"agentRef",
		&returns,
	)
	return returns
}


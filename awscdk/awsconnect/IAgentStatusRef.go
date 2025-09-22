package awsconnect

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsconnect/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentStatus.
// Experimental.
type IAgentStatusRef interface {
	constructs.IConstruct
	// A reference to a AgentStatus resource.
	// Experimental.
	AgentStatusRef() *AgentStatusReference
}

// The jsii proxy for IAgentStatusRef
type jsiiProxy_IAgentStatusRef struct {
	internal.Type__constructsIConstruct
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


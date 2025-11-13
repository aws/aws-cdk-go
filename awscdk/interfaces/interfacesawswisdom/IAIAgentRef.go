package interfacesawswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIAgent.
// Experimental.
type IAIAgentRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AIAgent resource.
	// Experimental.
	AiAgentRef() *AIAgentReference
}

// The jsii proxy for IAIAgentRef
type jsiiProxy_IAIAgentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IAIAgentRef) AiAgentRef() *AIAgentReference {
	var returns *AIAgentReference
	_jsii_.Get(
		j,
		"aiAgentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIAgentRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIAgentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


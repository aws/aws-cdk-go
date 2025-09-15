package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIAgent.
// Experimental.
type IAIAgentRef interface {
	constructs.IConstruct
	// A reference to a AIAgent resource.
	// Experimental.
	AiAgentRef() *AIAgentReference
}

// The jsii proxy for IAIAgentRef
type jsiiProxy_IAIAgentRef struct {
	internal.Type__constructsIConstruct
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


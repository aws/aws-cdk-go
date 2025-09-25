package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIAgentVersion.
// Experimental.
type IAIAgentVersionRef interface {
	constructs.IConstruct
	// A reference to a AIAgentVersion resource.
	// Experimental.
	AiAgentVersionRef() *AIAgentVersionReference
}

// The jsii proxy for IAIAgentVersionRef
type jsiiProxy_IAIAgentVersionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAIAgentVersionRef) AiAgentVersionRef() *AIAgentVersionReference {
	var returns *AIAgentVersionReference
	_jsii_.Get(
		j,
		"aiAgentVersionRef",
		&returns,
	)
	return returns
}


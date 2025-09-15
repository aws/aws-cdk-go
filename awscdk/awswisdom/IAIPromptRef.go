package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIPrompt.
// Experimental.
type IAIPromptRef interface {
	constructs.IConstruct
	// A reference to a AIPrompt resource.
	// Experimental.
	AiPromptRef() *AIPromptReference
}

// The jsii proxy for IAIPromptRef
type jsiiProxy_IAIPromptRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAIPromptRef) AiPromptRef() *AIPromptReference {
	var returns *AIPromptReference
	_jsii_.Get(
		j,
		"aiPromptRef",
		&returns,
	)
	return returns
}


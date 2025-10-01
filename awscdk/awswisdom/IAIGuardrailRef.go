package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIGuardrail.
// Experimental.
type IAIGuardrailRef interface {
	constructs.IConstruct
	// A reference to a AIGuardrail resource.
	// Experimental.
	AiGuardrailRef() *AIGuardrailReference
}

// The jsii proxy for IAIGuardrailRef
type jsiiProxy_IAIGuardrailRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IAIGuardrailRef) AiGuardrailRef() *AIGuardrailReference {
	var returns *AIGuardrailReference
	_jsii_.Get(
		j,
		"aiGuardrailRef",
		&returns,
	)
	return returns
}


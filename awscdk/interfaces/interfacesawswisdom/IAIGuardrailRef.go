package interfacesawswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIGuardrail.
// Experimental.
type IAIGuardrailRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AIGuardrail resource.
	// Experimental.
	AiGuardrailRef() *AIGuardrailReference
}

// The jsii proxy for IAIGuardrailRef
type jsiiProxy_IAIGuardrailRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IAIGuardrailRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIGuardrailRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


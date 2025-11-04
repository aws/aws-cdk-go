package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIGuardrailVersion.
// Experimental.
type IAIGuardrailVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AIGuardrailVersion resource.
	// Experimental.
	AiGuardrailVersionRef() *AIGuardrailVersionReference
}

// The jsii proxy for IAIGuardrailVersionRef
type jsiiProxy_IAIGuardrailVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAIGuardrailVersionRef) AiGuardrailVersionRef() *AIGuardrailVersionReference {
	var returns *AIGuardrailVersionReference
	_jsii_.Get(
		j,
		"aiGuardrailVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIGuardrailVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIGuardrailVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


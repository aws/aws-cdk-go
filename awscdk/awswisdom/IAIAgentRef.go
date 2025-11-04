package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIAgent.
// Experimental.
type IAIAgentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AIAgent resource.
	// Experimental.
	AiAgentRef() *AIAgentReference
}

// The jsii proxy for IAIAgentRef
type jsiiProxy_IAIAgentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAIAgentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
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


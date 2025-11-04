package awswisdom

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awswisdom/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AIAgentVersion.
// Experimental.
type IAIAgentVersionRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a AIAgentVersion resource.
	// Experimental.
	AiAgentVersionRef() *AIAgentVersionReference
}

// The jsii proxy for IAIAgentVersionRef
type jsiiProxy_IAIAgentVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IAIAgentVersionRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAIAgentVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


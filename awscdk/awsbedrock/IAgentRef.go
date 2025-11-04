package awsbedrock

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsbedrock/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Agent.
// Experimental.
type IAgentRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Agent resource.
	// Experimental.
	AgentRef() *AgentReference
}

// The jsii proxy for IAgentRef
type jsiiProxy_IAgentRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IAgentRef) AgentRef() *AgentReference {
	var returns *AgentReference
	_jsii_.Get(
		j,
		"agentRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


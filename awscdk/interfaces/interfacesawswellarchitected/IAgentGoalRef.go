package interfacesawswellarchitected

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawswellarchitected/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AgentGoal.
// Experimental.
type IAgentGoalRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a AgentGoal resource.
	// Experimental.
	AgentGoalRef() *AgentGoalReference
}

// The jsii proxy for IAgentGoalRef
type jsiiProxy_IAgentGoalRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IAgentGoalRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IAgentGoalRef) AgentGoalRef() *AgentGoalReference {
	var returns *AgentGoalReference
	_jsii_.Get(
		j,
		"agentGoalRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentGoalRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IAgentGoalRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


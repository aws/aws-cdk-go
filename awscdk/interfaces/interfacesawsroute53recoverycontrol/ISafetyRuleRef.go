package interfacesawsroute53recoverycontrol

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53recoverycontrol/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SafetyRule.
// Experimental.
type ISafetyRuleRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a SafetyRule resource.
	// Experimental.
	SafetyRuleRef() *SafetyRuleReference
}

// The jsii proxy for ISafetyRuleRef
type jsiiProxy_ISafetyRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ISafetyRuleRef) SafetyRuleRef() *SafetyRuleReference {
	var returns *SafetyRuleReference
	_jsii_.Get(
		j,
		"safetyRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISafetyRuleRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ISafetyRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


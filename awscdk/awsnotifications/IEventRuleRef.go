package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventRule.
// Experimental.
type IEventRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EventRule resource.
	// Experimental.
	EventRuleRef() *EventRuleReference
}

// The jsii proxy for IEventRuleRef
type jsiiProxy_IEventRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEventRuleRef) EventRuleRef() *EventRuleReference {
	var returns *EventRuleReference
	_jsii_.Get(
		j,
		"eventRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


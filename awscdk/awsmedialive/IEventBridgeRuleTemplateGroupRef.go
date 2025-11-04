package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBridgeRuleTemplateGroup.
// Experimental.
type IEventBridgeRuleTemplateGroupRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a EventBridgeRuleTemplateGroup resource.
	// Experimental.
	EventBridgeRuleTemplateGroupRef() *EventBridgeRuleTemplateGroupReference
}

// The jsii proxy for IEventBridgeRuleTemplateGroupRef
type jsiiProxy_IEventBridgeRuleTemplateGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_IEventBridgeRuleTemplateGroupRef) EventBridgeRuleTemplateGroupRef() *EventBridgeRuleTemplateGroupReference {
	var returns *EventBridgeRuleTemplateGroupReference
	_jsii_.Get(
		j,
		"eventBridgeRuleTemplateGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBridgeRuleTemplateGroupRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IEventBridgeRuleTemplateGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


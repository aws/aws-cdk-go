package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBridgeRuleTemplateGroup.
// Experimental.
type IEventBridgeRuleTemplateGroupRef interface {
	constructs.IConstruct
	// A reference to a EventBridgeRuleTemplateGroup resource.
	// Experimental.
	EventBridgeRuleTemplateGroupRef() *EventBridgeRuleTemplateGroupReference
}

// The jsii proxy for IEventBridgeRuleTemplateGroupRef
type jsiiProxy_IEventBridgeRuleTemplateGroupRef struct {
	internal.Type__constructsIConstruct
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


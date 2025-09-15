package awsmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventBridgeRuleTemplate.
// Experimental.
type IEventBridgeRuleTemplateRef interface {
	constructs.IConstruct
	// A reference to a EventBridgeRuleTemplate resource.
	// Experimental.
	EventBridgeRuleTemplateRef() *EventBridgeRuleTemplateReference
}

// The jsii proxy for IEventBridgeRuleTemplateRef
type jsiiProxy_IEventBridgeRuleTemplateRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IEventBridgeRuleTemplateRef) EventBridgeRuleTemplateRef() *EventBridgeRuleTemplateReference {
	var returns *EventBridgeRuleTemplateReference
	_jsii_.Get(
		j,
		"eventBridgeRuleTemplateRef",
		&returns,
	)
	return returns
}


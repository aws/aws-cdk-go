package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a EventRule.
// Experimental.
type IEventRuleRef interface {
	constructs.IConstruct
	// A reference to a EventRule resource.
	// Experimental.
	EventRuleRef() *EventRuleReference
}

// The jsii proxy for IEventRuleRef
type jsiiProxy_IEventRuleRef struct {
	internal.Type__constructsIConstruct
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


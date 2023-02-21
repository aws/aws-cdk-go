package awscodestarnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a notification target That allows AWS Chatbot and SNS topic to associate with this rule target.
type INotificationRuleTarget interface {
	// Returns a target configuration for notification rule.
	BindAsNotificationRuleTarget(scope constructs.Construct) *NotificationRuleTargetConfig
}

// The jsii proxy for INotificationRuleTarget
type jsiiProxy_INotificationRuleTarget struct {
	_ byte // padding
}

func (i *jsiiProxy_INotificationRuleTarget) BindAsNotificationRuleTarget(scope constructs.Construct) *NotificationRuleTargetConfig {
	if err := i.validateBindAsNotificationRuleTargetParameters(scope); err != nil {
		panic(err)
	}
	var returns *NotificationRuleTargetConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleTarget",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


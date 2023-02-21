package awscodestarnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a notification source The source that allows CodeBuild and CodePipeline to associate with this rule.
type INotificationRuleSource interface {
	// Returns a source configuration for notification rule.
	BindAsNotificationRuleSource(scope constructs.Construct) *NotificationRuleSourceConfig
}

// The jsii proxy for INotificationRuleSource
type jsiiProxy_INotificationRuleSource struct {
	_ byte // padding
}

func (i *jsiiProxy_INotificationRuleSource) BindAsNotificationRuleSource(scope constructs.Construct) *NotificationRuleSourceConfig {
	if err := i.validateBindAsNotificationRuleSourceParameters(scope); err != nil {
		panic(err)
	}
	var returns *NotificationRuleSourceConfig

	_jsii_.Invoke(
		i,
		"bindAsNotificationRuleSource",
		[]interface{}{scope},
		&returns,
	)

	return returns
}


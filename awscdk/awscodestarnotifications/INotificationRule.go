package awscodestarnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications/internal"
)

// Represents a notification rule.
type INotificationRule interface {
	awscdk.IResource
	// Adds target to notification rule.
	//
	// Returns: boolean - return true if it had any effect.
	AddTarget(target INotificationRuleTarget) *bool
	// The ARN of the notification rule (i.e. arn:aws:codestar-notifications:::notificationrule/01234abcde).
	NotificationRuleArn() *string
}

// The jsii proxy for INotificationRule
type jsiiProxy_INotificationRule struct {
	internal.Type__awscdkIResource
}

func (i *jsiiProxy_INotificationRule) AddTarget(target INotificationRuleTarget) *bool {
	if err := i.validateAddTargetParameters(target); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		i,
		"addTarget",
		[]interface{}{target},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_INotificationRule) NotificationRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationRuleArn",
		&returns,
	)
	return returns
}


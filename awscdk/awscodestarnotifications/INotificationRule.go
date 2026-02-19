package awscodestarnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications/internal"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscodestarnotifications"
	"github.com/aws/constructs-go/constructs/v10"
)

// Represents a notification rule.
type INotificationRule interface {
	interfacesawscodestarnotifications.INotificationRuleRef
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
	internal.Type__interfacesawscodestarnotificationsINotificationRuleRef
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

func (i *jsiiProxy_INotificationRule) ApplyRemovalPolicy(policy awscdk.RemovalPolicy) {
	if err := i.validateApplyRemovalPolicyParameters(policy); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"applyRemovalPolicy",
		[]interface{}{policy},
	)
}

func (i *jsiiProxy_INotificationRule) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_INotificationRule) NotificationRuleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"notificationRuleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationRule) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationRule) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationRule) NotificationRuleRef() *interfacesawscodestarnotifications.NotificationRuleReference {
	var returns *interfacesawscodestarnotifications.NotificationRuleReference
	_jsii_.Get(
		j,
		"notificationRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationRule) Stack() awscdk.Stack {
	var returns awscdk.Stack
	_jsii_.Get(
		j,
		"stack",
		&returns,
	)
	return returns
}


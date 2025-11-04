package awscodestarnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationRule.
// Experimental.
type INotificationRuleRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NotificationRule resource.
	// Experimental.
	NotificationRuleRef() *NotificationRuleReference
}

// The jsii proxy for INotificationRuleRef
type jsiiProxy_INotificationRuleRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INotificationRuleRef) NotificationRuleRef() *NotificationRuleReference {
	var returns *NotificationRuleReference
	_jsii_.Get(
		j,
		"notificationRuleRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationRuleRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationRuleRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


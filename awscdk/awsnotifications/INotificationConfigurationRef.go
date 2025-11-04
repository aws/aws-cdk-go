package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationConfiguration.
// Experimental.
type INotificationConfigurationRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NotificationConfiguration resource.
	// Experimental.
	NotificationConfigurationRef() *NotificationConfigurationReference
}

// The jsii proxy for INotificationConfigurationRef
type jsiiProxy_INotificationConfigurationRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INotificationConfigurationRef) NotificationConfigurationRef() *NotificationConfigurationReference {
	var returns *NotificationConfigurationReference
	_jsii_.Get(
		j,
		"notificationConfigurationRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationConfigurationRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationConfigurationRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


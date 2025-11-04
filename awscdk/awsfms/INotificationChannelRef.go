package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationChannel.
// Experimental.
type INotificationChannelRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a NotificationChannel resource.
	// Experimental.
	NotificationChannelRef() *NotificationChannelReference
}

// The jsii proxy for INotificationChannelRef
type jsiiProxy_INotificationChannelRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_INotificationChannelRef) NotificationChannelRef() *NotificationChannelReference {
	var returns *NotificationChannelReference
	_jsii_.Get(
		j,
		"notificationChannelRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationChannelRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INotificationChannelRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


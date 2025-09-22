package awsfms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationChannel.
// Experimental.
type INotificationChannelRef interface {
	constructs.IConstruct
	// A reference to a NotificationChannel resource.
	// Experimental.
	NotificationChannelRef() *NotificationChannelReference
}

// The jsii proxy for INotificationChannelRef
type jsiiProxy_INotificationChannelRef struct {
	internal.Type__constructsIConstruct
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


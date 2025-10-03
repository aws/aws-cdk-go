package awsfms

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfms/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationChannel.
// Experimental.
type INotificationChannelRef interface {
	constructs.IConstruct
}

// The jsii proxy for INotificationChannelRef
type jsiiProxy_INotificationChannelRef struct {
	internal.Type__constructsIConstruct
}


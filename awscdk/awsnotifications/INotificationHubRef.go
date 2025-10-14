package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationHub.
// Experimental.
type INotificationHubRef interface {
	constructs.IConstruct
	// A reference to a NotificationHub resource.
	// Experimental.
	NotificationHubRef() *NotificationHubReference
}

// The jsii proxy for INotificationHubRef
type jsiiProxy_INotificationHubRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INotificationHubRef) NotificationHubRef() *NotificationHubReference {
	var returns *NotificationHubReference
	_jsii_.Get(
		j,
		"notificationHubRef",
		&returns,
	)
	return returns
}


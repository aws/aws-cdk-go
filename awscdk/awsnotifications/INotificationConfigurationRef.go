package awsnotifications

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationConfiguration.
// Experimental.
type INotificationConfigurationRef interface {
	constructs.IConstruct
	// A reference to a NotificationConfiguration resource.
	// Experimental.
	NotificationConfigurationRef() *NotificationConfigurationReference
}

// The jsii proxy for INotificationConfigurationRef
type jsiiProxy_INotificationConfigurationRef struct {
	internal.Type__constructsIConstruct
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


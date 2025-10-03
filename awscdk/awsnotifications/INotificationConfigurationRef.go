package awsnotifications

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationConfiguration.
// Experimental.
type INotificationConfigurationRef interface {
	constructs.IConstruct
}

// The jsii proxy for INotificationConfigurationRef
type jsiiProxy_INotificationConfigurationRef struct {
	internal.Type__constructsIConstruct
}


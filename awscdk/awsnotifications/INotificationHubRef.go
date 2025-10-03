package awsnotifications

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationHub.
// Experimental.
type INotificationHubRef interface {
	constructs.IConstruct
}

// The jsii proxy for INotificationHubRef
type jsiiProxy_INotificationHubRef struct {
	internal.Type__constructsIConstruct
}


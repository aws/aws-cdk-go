package awscodestarnotifications

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarnotifications/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NotificationRule.
// Experimental.
type INotificationRuleRef interface {
	constructs.IConstruct
}

// The jsii proxy for INotificationRuleRef
type jsiiProxy_INotificationRuleRef struct {
	internal.Type__constructsIConstruct
}


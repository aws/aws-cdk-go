package awscloudfront

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MonitoringSubscription.
// Experimental.
type IMonitoringSubscriptionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMonitoringSubscriptionRef
type jsiiProxy_IMonitoringSubscriptionRef struct {
	internal.Type__constructsIConstruct
}


package awscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MonitoringSubscription.
// Experimental.
type IMonitoringSubscriptionRef interface {
	constructs.IConstruct
	// A reference to a MonitoringSubscription resource.
	// Experimental.
	MonitoringSubscriptionRef() *MonitoringSubscriptionReference
}

// The jsii proxy for IMonitoringSubscriptionRef
type jsiiProxy_IMonitoringSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IMonitoringSubscriptionRef) MonitoringSubscriptionRef() *MonitoringSubscriptionReference {
	var returns *MonitoringSubscriptionReference
	_jsii_.Get(
		j,
		"monitoringSubscriptionRef",
		&returns,
	)
	return returns
}


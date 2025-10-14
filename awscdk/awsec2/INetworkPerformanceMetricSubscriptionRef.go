package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkPerformanceMetricSubscription.
// Experimental.
type INetworkPerformanceMetricSubscriptionRef interface {
	constructs.IConstruct
	// A reference to a NetworkPerformanceMetricSubscription resource.
	// Experimental.
	NetworkPerformanceMetricSubscriptionRef() *NetworkPerformanceMetricSubscriptionReference
}

// The jsii proxy for INetworkPerformanceMetricSubscriptionRef
type jsiiProxy_INetworkPerformanceMetricSubscriptionRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_INetworkPerformanceMetricSubscriptionRef) NetworkPerformanceMetricSubscriptionRef() *NetworkPerformanceMetricSubscriptionReference {
	var returns *NetworkPerformanceMetricSubscriptionReference
	_jsii_.Get(
		j,
		"networkPerformanceMetricSubscriptionRef",
		&returns,
	)
	return returns
}


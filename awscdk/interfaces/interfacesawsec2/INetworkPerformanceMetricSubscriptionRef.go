package interfacesawsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsec2/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a NetworkPerformanceMetricSubscription.
// Experimental.
type INetworkPerformanceMetricSubscriptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a NetworkPerformanceMetricSubscription resource.
	// Experimental.
	NetworkPerformanceMetricSubscriptionRef() *NetworkPerformanceMetricSubscriptionReference
}

// The jsii proxy for INetworkPerformanceMetricSubscriptionRef
type jsiiProxy_INetworkPerformanceMetricSubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_INetworkPerformanceMetricSubscriptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_INetworkPerformanceMetricSubscriptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


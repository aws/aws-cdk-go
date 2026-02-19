package interfacesawscloudfront

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawscloudfront/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MonitoringSubscription.
// Experimental.
type IMonitoringSubscriptionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MonitoringSubscription resource.
	// Experimental.
	MonitoringSubscriptionRef() *MonitoringSubscriptionReference
}

// The jsii proxy for IMonitoringSubscriptionRef
type jsiiProxy_IMonitoringSubscriptionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IMonitoringSubscriptionRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IMonitoringSubscriptionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMonitoringSubscriptionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


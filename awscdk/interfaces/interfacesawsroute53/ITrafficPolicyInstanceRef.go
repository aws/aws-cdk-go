package interfacesawsroute53

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsroute53/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TrafficPolicyInstance.
// Experimental.
type ITrafficPolicyInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a TrafficPolicyInstance resource.
	// Experimental.
	TrafficPolicyInstanceRef() *TrafficPolicyInstanceReference
}

// The jsii proxy for ITrafficPolicyInstanceRef
type jsiiProxy_ITrafficPolicyInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_ITrafficPolicyInstanceRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
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

func (j *jsiiProxy_ITrafficPolicyInstanceRef) TrafficPolicyInstanceRef() *TrafficPolicyInstanceReference {
	var returns *TrafficPolicyInstanceReference
	_jsii_.Get(
		j,
		"trafficPolicyInstanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficPolicyInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ITrafficPolicyInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


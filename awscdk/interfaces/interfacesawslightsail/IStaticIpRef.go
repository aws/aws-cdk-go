package interfacesawslightsail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslightsail/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a StaticIp.
// Experimental.
type IStaticIpRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a StaticIp resource.
	// Experimental.
	StaticIpRef() *StaticIpReference
}

// The jsii proxy for IStaticIpRef
type jsiiProxy_IStaticIpRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IStaticIpRef) StaticIpRef() *StaticIpReference {
	var returns *StaticIpReference
	_jsii_.Get(
		j,
		"staticIpRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStaticIpRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IStaticIpRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


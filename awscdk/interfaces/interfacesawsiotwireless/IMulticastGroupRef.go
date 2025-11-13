package interfacesawsiotwireless

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsiotwireless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a MulticastGroup.
// Experimental.
type IMulticastGroupRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a MulticastGroup resource.
	// Experimental.
	MulticastGroupRef() *MulticastGroupReference
}

// The jsii proxy for IMulticastGroupRef
type jsiiProxy_IMulticastGroupRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMulticastGroupRef) MulticastGroupRef() *MulticastGroupReference {
	var returns *MulticastGroupReference
	_jsii_.Get(
		j,
		"multicastGroupRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMulticastGroupRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMulticastGroupRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


package interfacesawsb2bi

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsb2bi/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Capability.
// Experimental.
type ICapabilityRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Capability resource.
	// Experimental.
	CapabilityRef() *CapabilityReference
}

// The jsii proxy for ICapabilityRef
type jsiiProxy_ICapabilityRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ICapabilityRef) CapabilityRef() *CapabilityReference {
	var returns *CapabilityReference
	_jsii_.Get(
		j,
		"capabilityRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapabilityRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ICapabilityRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


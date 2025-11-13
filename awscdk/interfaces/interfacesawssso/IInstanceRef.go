package interfacesawssso

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawssso/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Instance.
// Experimental.
type IInstanceRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Instance resource.
	// Experimental.
	InstanceRef() *InstanceReference
}

// The jsii proxy for IInstanceRef
type jsiiProxy_IInstanceRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IInstanceRef) InstanceRef() *InstanceReference {
	var returns *InstanceReference
	_jsii_.Get(
		j,
		"instanceRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IInstanceRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


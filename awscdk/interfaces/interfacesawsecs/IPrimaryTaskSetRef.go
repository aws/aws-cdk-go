package interfacesawsecs

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrimaryTaskSet.
// Experimental.
type IPrimaryTaskSetRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a PrimaryTaskSet resource.
	// Experimental.
	PrimaryTaskSetRef() *PrimaryTaskSetReference
}

// The jsii proxy for IPrimaryTaskSetRef
type jsiiProxy_IPrimaryTaskSetRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IPrimaryTaskSetRef) PrimaryTaskSetRef() *PrimaryTaskSetReference {
	var returns *PrimaryTaskSetReference
	_jsii_.Get(
		j,
		"primaryTaskSetRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrimaryTaskSetRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IPrimaryTaskSetRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


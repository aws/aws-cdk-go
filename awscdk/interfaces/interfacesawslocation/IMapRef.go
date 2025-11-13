package interfacesawslocation

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslocation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Map.
// Experimental.
type IMapRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Map resource.
	// Experimental.
	MapRef() *MapReference
}

// The jsii proxy for IMapRef
type jsiiProxy_IMapRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_IMapRef) MapRef() *MapReference {
	var returns *MapReference
	_jsii_.Get(
		j,
		"mapRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMapRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IMapRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


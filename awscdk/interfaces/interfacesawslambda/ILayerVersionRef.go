package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LayerVersion.
// Experimental.
type ILayerVersionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LayerVersion resource.
	// Experimental.
	LayerVersionRef() *LayerVersionReference
}

// The jsii proxy for ILayerVersionRef
type jsiiProxy_ILayerVersionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILayerVersionRef) LayerVersionRef() *LayerVersionReference {
	var returns *LayerVersionReference
	_jsii_.Get(
		j,
		"layerVersionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


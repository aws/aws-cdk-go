package awsopsworks

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopsworks/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Layer.
// Experimental.
type ILayerRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a Layer resource.
	// Experimental.
	LayerRef() *LayerReference
}

// The jsii proxy for ILayerRef
type jsiiProxy_ILayerRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
}

func (j *jsiiProxy_ILayerRef) LayerRef() *LayerReference {
	var returns *LayerReference
	_jsii_.Get(
		j,
		"layerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


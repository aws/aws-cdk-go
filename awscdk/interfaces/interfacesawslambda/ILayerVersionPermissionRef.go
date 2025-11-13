package interfacesawslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LayerVersionPermission.
// Experimental.
type ILayerVersionPermissionRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a LayerVersionPermission resource.
	// Experimental.
	LayerVersionPermissionRef() *LayerVersionPermissionReference
}

// The jsii proxy for ILayerVersionPermissionRef
type jsiiProxy_ILayerVersionPermissionRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (j *jsiiProxy_ILayerVersionPermissionRef) LayerVersionPermissionRef() *LayerVersionPermissionReference {
	var returns *LayerVersionPermissionReference
	_jsii_.Get(
		j,
		"layerVersionPermissionRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersionPermissionRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ILayerVersionPermissionRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


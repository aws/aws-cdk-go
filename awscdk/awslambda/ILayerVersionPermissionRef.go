package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LayerVersionPermission.
// Experimental.
type ILayerVersionPermissionRef interface {
	constructs.IConstruct
	// A reference to a LayerVersionPermission resource.
	// Experimental.
	LayerVersionPermissionRef() *LayerVersionPermissionReference
}

// The jsii proxy for ILayerVersionPermissionRef
type jsiiProxy_ILayerVersionPermissionRef struct {
	internal.Type__constructsIConstruct
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


package awslambda

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LayerVersion.
// Experimental.
type ILayerVersionRef interface {
	constructs.IConstruct
	// A reference to a LayerVersion resource.
	// Experimental.
	LayerVersionRef() *LayerVersionReference
}

// The jsii proxy for ILayerVersionRef
type jsiiProxy_ILayerVersionRef struct {
	internal.Type__constructsIConstruct
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


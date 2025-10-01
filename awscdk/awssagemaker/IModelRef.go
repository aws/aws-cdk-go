package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Model.
// Experimental.
type IModelRef interface {
	constructs.IConstruct
	// A reference to a Model resource.
	// Experimental.
	ModelRef() *ModelReference
}

// The jsii proxy for IModelRef
type jsiiProxy_IModelRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModelRef) ModelRef() *ModelReference {
	var returns *ModelReference
	_jsii_.Get(
		j,
		"modelRef",
		&returns,
	)
	return returns
}


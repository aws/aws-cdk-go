package awssagemaker

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ModelCard.
// Experimental.
type IModelCardRef interface {
	constructs.IConstruct
	// A reference to a ModelCard resource.
	// Experimental.
	ModelCardRef() *ModelCardReference
}

// The jsii proxy for IModelCardRef
type jsiiProxy_IModelCardRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IModelCardRef) ModelCardRef() *ModelCardReference {
	var returns *ModelCardReference
	_jsii_.Get(
		j,
		"modelCardRef",
		&returns,
	)
	return returns
}


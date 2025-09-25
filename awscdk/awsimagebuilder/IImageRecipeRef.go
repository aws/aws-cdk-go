package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageRecipe.
// Experimental.
type IImageRecipeRef interface {
	constructs.IConstruct
	// A reference to a ImageRecipe resource.
	// Experimental.
	ImageRecipeRef() *ImageRecipeReference
}

// The jsii proxy for IImageRecipeRef
type jsiiProxy_IImageRecipeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IImageRecipeRef) ImageRecipeRef() *ImageRecipeReference {
	var returns *ImageRecipeReference
	_jsii_.Get(
		j,
		"imageRecipeRef",
		&returns,
	)
	return returns
}


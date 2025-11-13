package interfacesawsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageRecipe.
// Experimental.
type IImageRecipeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a ImageRecipe resource.
	// Experimental.
	ImageRecipeRef() *ImageRecipeReference
}

// The jsii proxy for IImageRecipeRef
type jsiiProxy_IImageRecipeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
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

func (j *jsiiProxy_IImageRecipeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageRecipeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


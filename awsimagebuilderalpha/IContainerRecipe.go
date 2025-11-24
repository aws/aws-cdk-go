package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An EC2 Image Builder Container Recipe.
// Experimental.
type IContainerRecipe interface {
	IRecipeBase
	// The ARN of the container recipe.
	// Experimental.
	ContainerRecipeArn() *string
	// The name of the container recipe.
	// Experimental.
	ContainerRecipeName() *string
	// The version of the container recipe.
	// Experimental.
	ContainerRecipeVersion() *string
}

// The jsii proxy for IContainerRecipe
type jsiiProxy_IContainerRecipe struct {
	jsiiProxy_IRecipeBase
}

func (j *jsiiProxy_IContainerRecipe) ContainerRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContainerRecipe) ContainerRecipeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContainerRecipe) ContainerRecipeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerRecipeVersion",
		&returns,
	)
	return returns
}


package awsimagebuilderalpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An EC2 Image Builder Image Recipe.
// Experimental.
type IImageRecipe interface {
	IRecipeBase
	// The ARN of the image recipe.
	// Experimental.
	ImageRecipeArn() *string
	// The name of the image recipe.
	// Experimental.
	ImageRecipeName() *string
	// The version of the image recipe.
	// Experimental.
	ImageRecipeVersion() *string
}

// The jsii proxy for IImageRecipe
type jsiiProxy_IImageRecipe struct {
	jsiiProxy_IRecipeBase
}

func (j *jsiiProxy_IImageRecipe) ImageRecipeArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageRecipe) ImageRecipeName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IImageRecipe) ImageRecipeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"imageRecipeVersion",
		&returns,
	)
	return returns
}


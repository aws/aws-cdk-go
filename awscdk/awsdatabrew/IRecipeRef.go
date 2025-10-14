package awsdatabrew

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Recipe.
// Experimental.
type IRecipeRef interface {
	constructs.IConstruct
	// A reference to a Recipe resource.
	// Experimental.
	RecipeRef() *RecipeReference
}

// The jsii proxy for IRecipeRef
type jsiiProxy_IRecipeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IRecipeRef) RecipeRef() *RecipeReference {
	var returns *RecipeReference
	_jsii_.Get(
		j,
		"recipeRef",
		&returns,
	)
	return returns
}


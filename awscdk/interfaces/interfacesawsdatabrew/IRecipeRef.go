package interfacesawsdatabrew

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Recipe.
// Experimental.
type IRecipeRef interface {
	constructs.IConstruct
	interfaces.IEnvironmentAware
	// A reference to a Recipe resource.
	// Experimental.
	RecipeRef() *RecipeReference
}

// The jsii proxy for IRecipeRef
type jsiiProxy_IRecipeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__interfacesIEnvironmentAware
}

func (i *jsiiProxy_IRecipeRef) With(mixins ...constructs.IMixin) constructs.IConstruct {
	args := []interface{}{}
	for _, a := range mixins {
		args = append(args, a)
	}

	var returns constructs.IConstruct

	_jsii_.Invoke(
		i,
		"with",
		args,
		&returns,
	)

	return returns
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

func (j *jsiiProxy_IRecipeRef) Env() *interfaces.ResourceEnvironment {
	var returns *interfaces.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IRecipeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


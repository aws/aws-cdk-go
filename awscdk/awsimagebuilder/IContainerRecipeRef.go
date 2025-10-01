package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerRecipe.
// Experimental.
type IContainerRecipeRef interface {
	constructs.IConstruct
	// A reference to a ContainerRecipe resource.
	// Experimental.
	ContainerRecipeRef() *ContainerRecipeReference
}

// The jsii proxy for IContainerRecipeRef
type jsiiProxy_IContainerRecipeRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IContainerRecipeRef) ContainerRecipeRef() *ContainerRecipeReference {
	var returns *ContainerRecipeReference
	_jsii_.Get(
		j,
		"containerRecipeRef",
		&returns,
	)
	return returns
}


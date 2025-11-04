package awsimagebuilder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerRecipe.
// Experimental.
type IContainerRecipeRef interface {
	constructs.IConstruct
	awscdk.IEnvironmentAware
	// A reference to a ContainerRecipe resource.
	// Experimental.
	ContainerRecipeRef() *ContainerRecipeReference
}

// The jsii proxy for IContainerRecipeRef
type jsiiProxy_IContainerRecipeRef struct {
	internal.Type__constructsIConstruct
	internal.Type__awscdkIEnvironmentAware
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

func (j *jsiiProxy_IContainerRecipeRef) Env() *awscdk.ResourceEnvironment {
	var returns *awscdk.ResourceEnvironment
	_jsii_.Get(
		j,
		"env",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IContainerRecipeRef) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}


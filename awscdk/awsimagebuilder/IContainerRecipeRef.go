package awsimagebuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ContainerRecipe.
// Experimental.
type IContainerRecipeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IContainerRecipeRef
type jsiiProxy_IContainerRecipeRef struct {
	internal.Type__constructsIConstruct
}


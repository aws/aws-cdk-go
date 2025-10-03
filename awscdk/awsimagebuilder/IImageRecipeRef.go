package awsimagebuilder

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsimagebuilder/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageRecipe.
// Experimental.
type IImageRecipeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IImageRecipeRef
type jsiiProxy_IImageRecipeRef struct {
	internal.Type__constructsIConstruct
}


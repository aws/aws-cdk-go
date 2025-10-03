package awsdatabrew

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatabrew/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Recipe.
// Experimental.
type IRecipeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRecipeRef
type jsiiProxy_IRecipeRef struct {
	internal.Type__constructsIConstruct
}


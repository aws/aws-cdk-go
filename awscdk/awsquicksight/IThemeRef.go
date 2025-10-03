package awsquicksight

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsquicksight/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Theme.
// Experimental.
type IThemeRef interface {
	constructs.IConstruct
}

// The jsii proxy for IThemeRef
type jsiiProxy_IThemeRef struct {
	internal.Type__constructsIConstruct
}


package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a AppBlockBuilder.
// Experimental.
type IAppBlockBuilderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IAppBlockBuilderRef
type jsiiProxy_IAppBlockBuilderRef struct {
	internal.Type__constructsIConstruct
}


package awsappstream

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsappstream/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageBuilder.
// Experimental.
type IImageBuilderRef interface {
	constructs.IConstruct
}

// The jsii proxy for IImageBuilderRef
type jsiiProxy_IImageBuilderRef struct {
	internal.Type__constructsIConstruct
}


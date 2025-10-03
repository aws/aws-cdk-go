package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ImageVersion.
// Experimental.
type IImageVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IImageVersionRef
type jsiiProxy_IImageVersionRef struct {
	internal.Type__constructsIConstruct
}


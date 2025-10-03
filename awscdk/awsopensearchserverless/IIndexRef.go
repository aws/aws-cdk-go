package awsopensearchserverless

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsopensearchserverless/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Index.
// Experimental.
type IIndexRef interface {
	constructs.IConstruct
}

// The jsii proxy for IIndexRef
type jsiiProxy_IIndexRef struct {
	internal.Type__constructsIConstruct
}


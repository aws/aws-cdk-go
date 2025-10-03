package awssagemaker

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssagemaker/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Space.
// Experimental.
type ISpaceRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISpaceRef
type jsiiProxy_ISpaceRef struct {
	internal.Type__constructsIConstruct
}


package awslakeformation

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslakeformation/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Tag.
// Experimental.
type ITagRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITagRef
type jsiiProxy_ITagRef struct {
	internal.Type__constructsIConstruct
}


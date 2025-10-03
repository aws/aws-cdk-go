package awslambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Version.
// Experimental.
type IVersionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVersionRef
type jsiiProxy_IVersionRef struct {
	internal.Type__constructsIConstruct
}


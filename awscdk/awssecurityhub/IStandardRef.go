package awssecurityhub

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awssecurityhub/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Standard.
// Experimental.
type IStandardRef interface {
	constructs.IConstruct
}

// The jsii proxy for IStandardRef
type jsiiProxy_IStandardRef struct {
	internal.Type__constructsIConstruct
}


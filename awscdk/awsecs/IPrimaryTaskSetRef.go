package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PrimaryTaskSet.
// Experimental.
type IPrimaryTaskSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPrimaryTaskSetRef
type jsiiProxy_IPrimaryTaskSetRef struct {
	internal.Type__constructsIConstruct
}


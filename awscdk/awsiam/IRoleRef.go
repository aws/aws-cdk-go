package awsiam

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Role.
// Experimental.
type IRoleRef interface {
	constructs.IConstruct
}

// The jsii proxy for IRoleRef
type jsiiProxy_IRoleRef struct {
	internal.Type__constructsIConstruct
}


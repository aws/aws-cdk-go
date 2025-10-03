package awscodestarconnections

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awscodestarconnections/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Connection.
// Experimental.
type IConnectionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IConnectionRef
type jsiiProxy_IConnectionRef struct {
	internal.Type__constructsIConstruct
}


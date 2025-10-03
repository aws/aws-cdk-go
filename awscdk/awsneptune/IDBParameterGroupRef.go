package awsneptune

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsneptune/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a DBParameterGroup.
// Experimental.
type IDBParameterGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDBParameterGroupRef
type jsiiProxy_IDBParameterGroupRef struct {
	internal.Type__constructsIConstruct
}


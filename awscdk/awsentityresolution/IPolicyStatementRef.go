package awsentityresolution

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsentityresolution/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PolicyStatement.
// Experimental.
type IPolicyStatementRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPolicyStatementRef
type jsiiProxy_IPolicyStatementRef struct {
	internal.Type__constructsIConstruct
}


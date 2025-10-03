package awsathena

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsathena/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a PreparedStatement.
// Experimental.
type IPreparedStatementRef interface {
	constructs.IConstruct
}

// The jsii proxy for IPreparedStatementRef
type jsiiProxy_IPreparedStatementRef struct {
	internal.Type__constructsIConstruct
}


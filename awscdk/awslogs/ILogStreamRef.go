package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogStream.
// Experimental.
type ILogStreamRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILogStreamRef
type jsiiProxy_ILogStreamRef struct {
	internal.Type__constructsIConstruct
}


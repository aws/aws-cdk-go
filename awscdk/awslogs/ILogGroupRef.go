package awslogs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslogs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a LogGroup.
// Experimental.
type ILogGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for ILogGroupRef
type jsiiProxy_ILogGroupRef struct {
	internal.Type__constructsIConstruct
}


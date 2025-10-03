package awsecs

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsecs/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a TaskSet.
// Experimental.
type ITaskSetRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITaskSetRef
type jsiiProxy_ITaskSetRef struct {
	internal.Type__constructsIConstruct
}


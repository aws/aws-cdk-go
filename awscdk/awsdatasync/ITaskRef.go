package awsdatasync

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdatasync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Task.
// Experimental.
type ITaskRef interface {
	constructs.IConstruct
}

// The jsii proxy for ITaskRef
type jsiiProxy_ITaskRef struct {
	internal.Type__constructsIConstruct
}


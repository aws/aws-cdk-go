package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a List.
// Experimental.
type IListRef interface {
	constructs.IConstruct
}

// The jsii proxy for IListRef
type jsiiProxy_IListRef struct {
	internal.Type__constructsIConstruct
}


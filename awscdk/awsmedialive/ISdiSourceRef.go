package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a SdiSource.
// Experimental.
type ISdiSourceRef interface {
	constructs.IConstruct
}

// The jsii proxy for ISdiSourceRef
type jsiiProxy_ISdiSourceRef struct {
	internal.Type__constructsIConstruct
}


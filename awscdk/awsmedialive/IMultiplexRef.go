package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Multiplex.
// Experimental.
type IMultiplexRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMultiplexRef
type jsiiProxy_IMultiplexRef struct {
	internal.Type__constructsIConstruct
}


package awsmedialive

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsmedialive/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Multiplexprogram.
// Experimental.
type IMultiplexprogramRef interface {
	constructs.IConstruct
}

// The jsii proxy for IMultiplexprogramRef
type jsiiProxy_IMultiplexprogramRef struct {
	internal.Type__constructsIConstruct
}


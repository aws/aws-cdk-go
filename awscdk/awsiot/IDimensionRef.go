package awsiot

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiot/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Dimension.
// Experimental.
type IDimensionRef interface {
	constructs.IConstruct
}

// The jsii proxy for IDimensionRef
type jsiiProxy_IDimensionRef struct {
	internal.Type__constructsIConstruct
}


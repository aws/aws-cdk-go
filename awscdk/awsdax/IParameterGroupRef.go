package awsdax

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdax/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ParameterGroup.
// Experimental.
type IParameterGroupRef interface {
	constructs.IConstruct
}

// The jsii proxy for IParameterGroupRef
type jsiiProxy_IParameterGroupRef struct {
	internal.Type__constructsIConstruct
}


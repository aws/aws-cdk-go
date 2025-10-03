package awsssm

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Parameter.
// Experimental.
type IParameterRef interface {
	constructs.IConstruct
}

// The jsii proxy for IParameterRef
type jsiiProxy_IParameterRef struct {
	internal.Type__constructsIConstruct
}


package awsfrauddetector

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Variable.
// Experimental.
type IVariableRef interface {
	constructs.IConstruct
}

// The jsii proxy for IVariableRef
type jsiiProxy_IVariableRef struct {
	internal.Type__constructsIConstruct
}


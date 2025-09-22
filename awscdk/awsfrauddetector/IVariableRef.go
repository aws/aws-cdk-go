package awsfrauddetector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsfrauddetector/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Variable.
// Experimental.
type IVariableRef interface {
	constructs.IConstruct
	// A reference to a Variable resource.
	// Experimental.
	VariableRef() *VariableReference
}

// The jsii proxy for IVariableRef
type jsiiProxy_IVariableRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IVariableRef) VariableRef() *VariableReference {
	var returns *VariableReference
	_jsii_.Get(
		j,
		"variableRef",
		&returns,
	)
	return returns
}


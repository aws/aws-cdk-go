package awsmemorydb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsmemorydb/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a ParameterGroup.
// Experimental.
type IParameterGroupRef interface {
	constructs.IConstruct
	// A reference to a ParameterGroup resource.
	// Experimental.
	ParameterGroupRef() *ParameterGroupReference
}

// The jsii proxy for IParameterGroupRef
type jsiiProxy_IParameterGroupRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IParameterGroupRef) ParameterGroupRef() *ParameterGroupReference {
	var returns *ParameterGroupReference
	_jsii_.Get(
		j,
		"parameterGroupRef",
		&returns,
	)
	return returns
}


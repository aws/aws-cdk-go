package awsssm

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsssm/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a Parameter.
// Experimental.
type IParameterRef interface {
	constructs.IConstruct
	// A reference to a Parameter resource.
	// Experimental.
	ParameterRef() *ParameterReference
}

// The jsii proxy for IParameterRef
type jsiiProxy_IParameterRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IParameterRef) ParameterRef() *ParameterReference {
	var returns *ParameterReference
	_jsii_.Get(
		j,
		"parameterRef",
		&returns,
	)
	return returns
}


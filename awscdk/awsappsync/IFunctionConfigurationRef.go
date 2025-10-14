package awsappsync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsappsync/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Indicates that this resource can be referenced as a FunctionConfiguration.
// Experimental.
type IFunctionConfigurationRef interface {
	constructs.IConstruct
	// A reference to a FunctionConfiguration resource.
	// Experimental.
	FunctionConfigurationRef() *FunctionConfigurationReference
}

// The jsii proxy for IFunctionConfigurationRef
type jsiiProxy_IFunctionConfigurationRef struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IFunctionConfigurationRef) FunctionConfigurationRef() *FunctionConfigurationReference {
	var returns *FunctionConfigurationReference
	_jsii_.Get(
		j,
		"functionConfigurationRef",
		&returns,
	)
	return returns
}


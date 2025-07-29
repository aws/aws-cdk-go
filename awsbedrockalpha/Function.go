package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a function in a function schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   function_ := bedrock_alpha.NewFunction(&FunctionProps{
//   	Description: jsii.String("description"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Parameters: map[string]functionParameterProps{
//   		"parametersKey": &functionParameterProps{
//   			"type": bedrock_alpha.ParameterType_STRING,
//
//   			// the properties below are optional
//   			"description": jsii.String("description"),
//   			"required": jsii.Boolean(false),
//   		},
//   	},
//   	RequireConfirmation: bedrock_alpha.RequireConfirmation_ENABLED,
//   })
//
// Experimental.
type Function interface {
	// Description of the function.
	// Experimental.
	Description() *string
	// The name of the function.
	// Experimental.
	Name() *string
	// Parameters for the function.
	// Experimental.
	Parameters() *map[string]FunctionParameter
	// Whether to require confirmation before executing the function.
	// Experimental.
	RequireConfirmation() RequireConfirmation
}

// The jsii proxy struct for Function
type jsiiProxy_Function struct {
	_ byte // padding
}

func (j *jsiiProxy_Function) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) Parameters() *map[string]FunctionParameter {
	var returns *map[string]FunctionParameter
	_jsii_.Get(
		j,
		"parameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Function) RequireConfirmation() RequireConfirmation {
	var returns RequireConfirmation
	_jsii_.Get(
		j,
		"requireConfirmation",
		&returns,
	)
	return returns
}


// Experimental.
func NewFunction(props *FunctionProps) Function {
	_init_.Initialize()

	if err := validateNewFunctionParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_Function{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Function",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewFunction_Override(f Function, props *FunctionProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.Function",
		[]interface{}{props},
		f,
	)
}


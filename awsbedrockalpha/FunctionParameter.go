package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a function parameter in a function schema.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   functionParameter := bedrock_alpha.NewFunctionParameter(&FunctionParameterProps{
//   	Type: bedrock_alpha.ParameterType_STRING,
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	Required: jsii.Boolean(false),
//   })
//
// Experimental.
type FunctionParameter interface {
	// Description of the parameter.
	// Default: undefined no description will be present.
	//
	// Experimental.
	Description() *string
	// Whether the parameter is required.
	// Experimental.
	Required() *bool
	// The type of the parameter.
	// Experimental.
	Type() ParameterType
}

// The jsii proxy struct for FunctionParameter
type jsiiProxy_FunctionParameter struct {
	_ byte // padding
}

func (j *jsiiProxy_FunctionParameter) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionParameter) Required() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"required",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionParameter) Type() ParameterType {
	var returns ParameterType
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Experimental.
func NewFunctionParameter(props *FunctionParameterProps) FunctionParameter {
	_init_.Initialize()

	if err := validateNewFunctionParameterParameters(props); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionParameter{}

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.FunctionParameter",
		[]interface{}{props},
		&j,
	)

	return &j
}

// Experimental.
func NewFunctionParameter_Override(f FunctionParameter, props *FunctionParameterProps) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.FunctionParameter",
		[]interface{}{props},
		f,
	)
}


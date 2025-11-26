package awsimagebuilderalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsimagebuilderalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The parameter value for a workflow parameter.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import imagebuilder_alpha "github.com/aws/aws-cdk-go/awsimagebuilderalpha"
//
//   workflowParameterValue := imagebuilder_alpha.WorkflowParameterValue_FromBoolean(jsii.Boolean(false))
//
// Experimental.
type WorkflowParameterValue interface {
	// The rendered parameter value.
	// Experimental.
	Value() *[]*string
}

// The jsii proxy struct for WorkflowParameterValue
type jsiiProxy_WorkflowParameterValue struct {
	_ byte // padding
}

func (j *jsiiProxy_WorkflowParameterValue) Value() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Experimental.
func NewWorkflowParameterValue(value *[]*string) WorkflowParameterValue {
	_init_.Initialize()

	if err := validateNewWorkflowParameterValueParameters(value); err != nil {
		panic(err)
	}
	j := jsiiProxy_WorkflowParameterValue{}

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		[]interface{}{value},
		&j,
	)

	return &j
}

// Experimental.
func NewWorkflowParameterValue_Override(w WorkflowParameterValue, value *[]*string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		[]interface{}{value},
		w,
	)
}

// The value of the parameter as a boolean.
// Experimental.
func WorkflowParameterValue_FromBoolean(value *bool) WorkflowParameterValue {
	_init_.Initialize()

	if err := validateWorkflowParameterValue_FromBooleanParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowParameterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		"fromBoolean",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The value of the parameter as an integer.
// Experimental.
func WorkflowParameterValue_FromInteger(value *float64) WorkflowParameterValue {
	_init_.Initialize()

	if err := validateWorkflowParameterValue_FromIntegerParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowParameterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		"fromInteger",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The value of the parameter as a string.
// Experimental.
func WorkflowParameterValue_FromString(value *string) WorkflowParameterValue {
	_init_.Initialize()

	if err := validateWorkflowParameterValue_FromStringParameters(value); err != nil {
		panic(err)
	}
	var returns WorkflowParameterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		"fromString",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// The value of the parameter as a string list.
// Experimental.
func WorkflowParameterValue_FromStringList(values *[]*string) WorkflowParameterValue {
	_init_.Initialize()

	if err := validateWorkflowParameterValue_FromStringListParameters(values); err != nil {
		panic(err)
	}
	var returns WorkflowParameterValue

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-imagebuilder-alpha.WorkflowParameterValue",
		"fromStringList",
		[]interface{}{values},
		&returns,
	)

	return returns
}


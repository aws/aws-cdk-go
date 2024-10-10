package awscdkpipesalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkpipesalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Transform or replace the input event payload.
//
// Example:
//   var sourceQueue queue
//   var targetQueue queue
//
//
//   pipeTarget := targets.NewSqsTarget(targetQueue, &SqsTargetParameters{
//   	InputTransformation: pipes.InputTransformation_FromObject(map[string]interface{}{
//   		"SomeKey": pipes.DynamicInput_fromEventPath(jsii.String("$.body")),
//   	}),
//   })
//
//   pipe := pipes.NewPipe(this, jsii.String("Pipe"), &PipeProps{
//   	Source: awscdkpipessourcesalpha.NewSqsSource(sourceQueue),
//   	Target: pipeTarget,
//   })
//
// Experimental.
type InputTransformation interface {
	IInputTransformation
	// Bind the input transformation to the pipe and returns the inputTemplate string.
	// Experimental.
	Bind(pipe IPipe) *InputTransformationConfig
}

// The jsii proxy struct for InputTransformation
type jsiiProxy_InputTransformation struct {
	jsiiProxy_IInputTransformation
}

// Creates an InputTransformation from a jsonPath expression of the input event.
// Experimental.
func InputTransformation_FromEventPath(jsonPathExpression *string) InputTransformation {
	_init_.Initialize()

	if err := validateInputTransformation_FromEventPathParameters(jsonPathExpression); err != nil {
		panic(err)
	}
	var returns InputTransformation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.InputTransformation",
		"fromEventPath",
		[]interface{}{jsonPathExpression},
		&returns,
	)

	return returns
}

// Creates an InputTransformation from a pipe variable.
// Experimental.
func InputTransformation_FromObject(inputTemplate *map[string]interface{}) InputTransformation {
	_init_.Initialize()

	if err := validateInputTransformation_FromObjectParameters(inputTemplate); err != nil {
		panic(err)
	}
	var returns InputTransformation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.InputTransformation",
		"fromObject",
		[]interface{}{inputTemplate},
		&returns,
	)

	return returns
}

// Creates an InputTransformation from a string.
// Experimental.
func InputTransformation_FromText(inputTemplate *string) InputTransformation {
	_init_.Initialize()

	if err := validateInputTransformation_FromTextParameters(inputTemplate); err != nil {
		panic(err)
	}
	var returns InputTransformation

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-pipes-alpha.InputTransformation",
		"fromText",
		[]interface{}{inputTemplate},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_InputTransformation) Bind(pipe IPipe) *InputTransformationConfig {
	if err := i.validateBindParameters(pipe); err != nil {
		panic(err)
	}
	var returns *InputTransformationConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{pipe},
		&returns,
	)

	return returns
}


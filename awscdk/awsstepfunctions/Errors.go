package awsstepfunctions

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Predefined error strings Error names in Amazon States Language - https://states-language.net/spec.html#appendix-a Error handling in Step Functions - https://docs.aws.amazon.com/step-functions/latest/dg/concepts-error-handling.html.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   errors := awscdk.Aws_stepfunctions.NewErrors()
//
type Errors interface {
}

// The jsii proxy struct for Errors
type jsiiProxy_Errors struct {
	_ byte // padding
}

func NewErrors() Errors {
	_init_.Initialize()

	j := jsiiProxy_Errors{}

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewErrors_Override(e Errors) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		nil, // no parameters
		e,
	)
}

func Errors_ALL() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"ALL",
		&returns,
	)
	return returns
}

func Errors_BRANCH_FAILED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"BRANCH_FAILED",
		&returns,
	)
	return returns
}

func Errors_HEARTBEAT_TIMEOUT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"HEARTBEAT_TIMEOUT",
		&returns,
	)
	return returns
}

func Errors_NO_CHOICE_MATCHED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"NO_CHOICE_MATCHED",
		&returns,
	)
	return returns
}

func Errors_PARAMETER_PATH_FAILURE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"PARAMETER_PATH_FAILURE",
		&returns,
	)
	return returns
}

func Errors_PERMISSIONS() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"PERMISSIONS",
		&returns,
	)
	return returns
}

func Errors_RESULT_PATH_MATCH_FAILURE() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"RESULT_PATH_MATCH_FAILURE",
		&returns,
	)
	return returns
}

func Errors_TASKS_FAILED() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"TASKS_FAILED",
		&returns,
	)
	return returns
}

func Errors_TIMEOUT() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions.Errors",
		"TIMEOUT",
		&returns,
	)
	return returns
}


package awslambda

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Specify the runtime update mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   runtimeManagementMode := awscdk.Aws_lambda.RuntimeManagementMode_Manual(jsii.String("arn"))
//
type RuntimeManagementMode interface {
	Arn() *string
	Mode() *string
	// https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-runtimemanagementconfig.html.
	RuntimeManagementConfig() *CfnFunction_RuntimeManagementConfigProperty
}

// The jsii proxy struct for RuntimeManagementMode
type jsiiProxy_RuntimeManagementMode struct {
	_ byte // padding
}

func (j *jsiiProxy_RuntimeManagementMode) Arn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"arn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuntimeManagementMode) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RuntimeManagementMode) RuntimeManagementConfig() *CfnFunction_RuntimeManagementConfigProperty {
	var returns *CfnFunction_RuntimeManagementConfigProperty
	_jsii_.Get(
		j,
		"runtimeManagementConfig",
		&returns,
	)
	return returns
}


func NewRuntimeManagementMode(mode *string, arn *string) RuntimeManagementMode {
	_init_.Initialize()

	if err := validateNewRuntimeManagementModeParameters(mode); err != nil {
		panic(err)
	}
	j := jsiiProxy_RuntimeManagementMode{}

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.RuntimeManagementMode",
		[]interface{}{mode, arn},
		&j,
	)

	return &j
}

func NewRuntimeManagementMode_Override(r RuntimeManagementMode, mode *string, arn *string) {
	_init_.Initialize()

	_jsii_.Create(
		"aws-cdk-lib.aws_lambda.RuntimeManagementMode",
		[]interface{}{mode, arn},
		r,
	)
}

// You specify a runtime version in your function configuration.
//
// The function uses this runtime version indefinitely.
// In the rare case in which a new runtime version is incompatible with an existing function,
// you can use this mode to roll back your function to an earlier runtime version.
func RuntimeManagementMode_Manual(arn *string) RuntimeManagementMode {
	_init_.Initialize()

	if err := validateRuntimeManagementMode_ManualParameters(arn); err != nil {
		panic(err)
	}
	var returns RuntimeManagementMode

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_lambda.RuntimeManagementMode",
		"manual",
		[]interface{}{arn},
		&returns,
	)

	return returns
}

func RuntimeManagementMode_AUTO() RuntimeManagementMode {
	_init_.Initialize()
	var returns RuntimeManagementMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.RuntimeManagementMode",
		"AUTO",
		&returns,
	)
	return returns
}

func RuntimeManagementMode_FUNCTION_UPDATE() RuntimeManagementMode {
	_init_.Initialize()
	var returns RuntimeManagementMode
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_lambda.RuntimeManagementMode",
		"FUNCTION_UPDATE",
		&returns,
	)
	return returns
}


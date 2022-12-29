package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// The generation of Elastic Inference (EI) instance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorClass := awscdk.Aws_stepfunctions_tasks.acceleratorClass.of(jsii.String("version"))
//
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html
//
type AcceleratorClass interface {
	// - Elastic Inference accelerator generation.
	Version() *string
}

// The jsii proxy struct for AcceleratorClass
type jsiiProxy_AcceleratorClass struct {
	_ byte // padding
}

func (j *jsiiProxy_AcceleratorClass) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}


// Custom AcceleratorType.
func AcceleratorClass_Of(version *string) AcceleratorClass {
	_init_.Initialize()

	if err := validateAcceleratorClass_OfParameters(version); err != nil {
		panic(err)
	}
	var returns AcceleratorClass

	_jsii_.StaticInvoke(
		"aws-cdk-lib.aws_stepfunctions_tasks.AcceleratorClass",
		"of",
		[]interface{}{version},
		&returns,
	)

	return returns
}

func AcceleratorClass_EIA1() AcceleratorClass {
	_init_.Initialize()
	var returns AcceleratorClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.AcceleratorClass",
		"EIA1",
		&returns,
	)
	return returns
}

func AcceleratorClass_EIA2() AcceleratorClass {
	_init_.Initialize()
	var returns AcceleratorClass
	_jsii_.StaticGet(
		"aws-cdk-lib.aws_stepfunctions_tasks.AcceleratorClass",
		"EIA2",
		&returns,
	)
	return returns
}


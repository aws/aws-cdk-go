package awsstepfunctionstasks

import (
	_init_ "github.com/aws/aws-cdk-go/awscdk/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/awsec2"
)

// The size of the Elastic Inference (EI) instance to use for the production variant.
//
// EI instances provide on-demand GPU computing for inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   acceleratorType := awscdk.Aws_stepfunctions_tasks.NewAcceleratorType(jsii.String("instanceTypeIdentifier"))
//
// See: https://docs.aws.amazon.com/sagemaker/latest/dg/ei.html
//
// Experimental.
type AcceleratorType interface {
	// Return the accelerator type as a dotted string.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcceleratorType
type jsiiProxy_AcceleratorType struct {
	_ byte // padding
}

// Experimental.
func NewAcceleratorType(instanceTypeIdentifier *string) AcceleratorType {
	_init_.Initialize()

	j := jsiiProxy_AcceleratorType{}

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AcceleratorType",
		[]interface{}{instanceTypeIdentifier},
		&j,
	)

	return &j
}

// Experimental.
func NewAcceleratorType_Override(a AcceleratorType, instanceTypeIdentifier *string) {
	_init_.Initialize()

	_jsii_.Create(
		"monocdk.aws_stepfunctions_tasks.AcceleratorType",
		[]interface{}{instanceTypeIdentifier},
		a,
	)
}

// AcceleratorType.
//
// This class takes a combination of a class and size.
// Experimental.
func AcceleratorType_Of(acceleratorClass AcceleratorClass, instanceSize awsec2.InstanceSize) AcceleratorType {
	_init_.Initialize()

	var returns AcceleratorType

	_jsii_.StaticInvoke(
		"monocdk.aws_stepfunctions_tasks.AcceleratorType",
		"of",
		[]interface{}{acceleratorClass, instanceSize},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AcceleratorType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Supported Elastic Inference (EI) instance types for SageMaker instance-based production variants.
//
// EI instances provide on-demand GPU computing for inference.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   acceleratorType := sagemaker_alpha.AcceleratorType_EIA1_LARGE()
//
// Experimental.
type AcceleratorType interface {
	// Return the accelerator type as a string.
	//
	// Returns: The accelerator type as a string.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AcceleratorType
type jsiiProxy_AcceleratorType struct {
	_ byte // padding
}

// Experimental.
func NewAcceleratorType(acceleratorType *string) AcceleratorType {
	_init_.Initialize()

	if err := validateNewAcceleratorTypeParameters(acceleratorType); err != nil {
		panic(err)
	}
	j := jsiiProxy_AcceleratorType{}

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		[]interface{}{acceleratorType},
		&j,
	)

	return &j
}

// Experimental.
func NewAcceleratorType_Override(a AcceleratorType, acceleratorType *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		[]interface{}{acceleratorType},
		a,
	)
}

// Builds an AcceleratorType from a given string or token (such as a CfnParameter).
//
// Returns: A strongly typed AcceleratorType.
// Experimental.
func AcceleratorType_Of(acceleratorType *string) AcceleratorType {
	_init_.Initialize()

	if err := validateAcceleratorType_OfParameters(acceleratorType); err != nil {
		panic(err)
	}
	var returns AcceleratorType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"of",
		[]interface{}{acceleratorType},
		&returns,
	)

	return returns
}

func AcceleratorType_EIA1_LARGE() AcceleratorType {
	_init_.Initialize()
	var returns AcceleratorType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"EIA1_LARGE",
		&returns,
	)
	return returns
}

func AcceleratorType_EIA1_MEDIUM() AcceleratorType {
	_init_.Initialize()
	var returns AcceleratorType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"EIA1_MEDIUM",
		&returns,
	)
	return returns
}

func AcceleratorType_EIA1_XLARGE() AcceleratorType {
	_init_.Initialize()
	var returns AcceleratorType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"EIA1_XLARGE",
		&returns,
	)
	return returns
}

func AcceleratorType_EIA2_LARGE() AcceleratorType {
	_init_.Initialize()
	var returns AcceleratorType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"EIA2_LARGE",
		&returns,
	)
	return returns
}

func AcceleratorType_EIA2_MEDIUM() AcceleratorType {
	_init_.Initialize()
	var returns AcceleratorType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"EIA2_MEDIUM",
		&returns,
	)
	return returns
}

func AcceleratorType_EIA2_XLARGE() AcceleratorType {
	_init_.Initialize()
	var returns AcceleratorType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.AcceleratorType",
		"EIA2_XLARGE",
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


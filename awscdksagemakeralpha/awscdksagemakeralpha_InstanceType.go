// The CDK Construct Library for AWS::SageMaker
package awscdksagemakeralpha

import (
	_init_ "github.com/aws/aws-cdk-go/awscdksagemakeralpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Supported instance types for SageMaker instance-based production variants.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import sagemaker_alpha "github.com/aws/aws-cdk-go/awscdksagemakeralpha"
//
//   instanceType := sagemaker_alpha.instanceType_C4_2XLARGE()
//
// Experimental.
type InstanceType interface {
	// Return the instance type as a string.
	//
	// Returns: The instance type as a string.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for InstanceType
type jsiiProxy_InstanceType struct {
	_ byte // padding
}

// Experimental.
func NewInstanceType(instanceType *string) InstanceType {
	_init_.Initialize()

	if err := validateNewInstanceTypeParameters(instanceType); err != nil {
		panic(err)
	}
	j := jsiiProxy_InstanceType{}

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		[]interface{}{instanceType},
		&j,
	)

	return &j
}

// Experimental.
func NewInstanceType_Override(i InstanceType, instanceType *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		[]interface{}{instanceType},
		i,
	)
}

// Builds an InstanceType from a given string or token (such as a CfnParameter).
//
// Returns: A strongly typed InstanceType.
// Experimental.
func InstanceType_Of(instanceType *string) InstanceType {
	_init_.Initialize()

	if err := validateInstanceType_OfParameters(instanceType); err != nil {
		panic(err)
	}
	var returns InstanceType

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"of",
		[]interface{}{instanceType},
		&returns,
	)

	return returns
}

func InstanceType_C4_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C4_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C4_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C4_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C4_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C4_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C4_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C4_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_C4_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C4_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5_18XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5_18XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5_9XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5_9XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5D_18XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5D_18XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5D_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5D_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5D_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5D_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5D_9XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5D_9XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5D_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5D_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_C5D_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C5D_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_32XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_32XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_C6I_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"C6I_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G4DN_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G4DN_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G4DN_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G4DN_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G4DN_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G4DN_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G4DN_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G4DN_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G4DN_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G4DN_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G4DN_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G4DN_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_48XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_48XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_G5_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"G5_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_INF1_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"INF1_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_INF1_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"INF1_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_INF1_6XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"INF1_6XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_INF1_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"INF1_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M4_10XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M4_10XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M4_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M4_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M4_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M4_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M4_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M4_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M4_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M4_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5D_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5D_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5D_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5D_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5D_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5D_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5D_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5D_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5D_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5D_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_M5D_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"M5D_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P2_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P2_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P2_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P2_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P2_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P2_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P3_16XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P3_16XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P3_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P3_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P3_8XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P3_8XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_P4D_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"P4D_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_12XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5D_12XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_24XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5D_24XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5D_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_4XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5D_4XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5D_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_R5D_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"R5D_XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_T2_2XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"T2_2XLARGE",
		&returns,
	)
	return returns
}

func InstanceType_T2_LARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"T2_LARGE",
		&returns,
	)
	return returns
}

func InstanceType_T2_MEDIUM() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"T2_MEDIUM",
		&returns,
	)
	return returns
}

func InstanceType_T2_XLARGE() InstanceType {
	_init_.Initialize()
	var returns InstanceType
	_jsii_.StaticGet(
		"@aws-cdk/aws-sagemaker-alpha.InstanceType",
		"T2_XLARGE",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InstanceType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


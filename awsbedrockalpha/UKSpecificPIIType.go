package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Types of PII specific to the United Kingdom (UK).
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   uKSpecificPIIType := bedrock_alpha.UKSpecificPIIType_UK_NATIONAL_HEALTH_SERVICE_NUMBER()
//
// Experimental.
type UKSpecificPIIType interface {
	PIIType
	// The string value of the PII type.
	// Experimental.
	Value() *string
	// Returns the string representation of the PII type.
	//
	// Returns: The string value of the PII type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for UKSpecificPIIType
type jsiiProxy_UKSpecificPIIType struct {
	jsiiProxy_PIIType
}

func (j *jsiiProxy_UKSpecificPIIType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func UKSpecificPIIType_UK_NATIONAL_HEALTH_SERVICE_NUMBER() UKSpecificPIIType {
	_init_.Initialize()
	var returns UKSpecificPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.UKSpecificPIIType",
		"UK_NATIONAL_HEALTH_SERVICE_NUMBER",
		&returns,
	)
	return returns
}

func UKSpecificPIIType_UK_NATIONAL_INSURANCE_NUMBER() UKSpecificPIIType {
	_init_.Initialize()
	var returns UKSpecificPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.UKSpecificPIIType",
		"UK_NATIONAL_INSURANCE_NUMBER",
		&returns,
	)
	return returns
}

func UKSpecificPIIType_UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER() UKSpecificPIIType {
	_init_.Initialize()
	var returns UKSpecificPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.UKSpecificPIIType",
		"UK_UNIQUE_TAXPAYER_REFERENCE_NUMBER",
		&returns,
	)
	return returns
}

func (u *jsiiProxy_UKSpecificPIIType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		u,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


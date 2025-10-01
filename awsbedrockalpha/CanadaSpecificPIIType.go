package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Types of PII specific to Canada.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   canadaSpecificPIIType := bedrock_alpha.CanadaSpecificPIIType_CA_HEALTH_NUMBER()
//
// Experimental.
type CanadaSpecificPIIType interface {
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

// The jsii proxy struct for CanadaSpecificPIIType
type jsiiProxy_CanadaSpecificPIIType struct {
	jsiiProxy_PIIType
}

func (j *jsiiProxy_CanadaSpecificPIIType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func CanadaSpecificPIIType_CA_HEALTH_NUMBER() CanadaSpecificPIIType {
	_init_.Initialize()
	var returns CanadaSpecificPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.CanadaSpecificPIIType",
		"CA_HEALTH_NUMBER",
		&returns,
	)
	return returns
}

func CanadaSpecificPIIType_CA_SOCIAL_INSURANCE_NUMBER() CanadaSpecificPIIType {
	_init_.Initialize()
	var returns CanadaSpecificPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.CanadaSpecificPIIType",
		"CA_SOCIAL_INSURANCE_NUMBER",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CanadaSpecificPIIType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


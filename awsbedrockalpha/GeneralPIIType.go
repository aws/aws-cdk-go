package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Types of PII that are general, and not domain-specific.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   })
//
//   // Add PII filter for addresses with input/output actions
//   guardrail.AddPIIFilter(&PIIFilter{
//   	Type: bedrock.GeneralPIIType_ADDRESS(),
//   	Action: bedrock.GuardrailAction_BLOCK,
//   	// below props are optional
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
//   	OutputEnabled: jsii.Boolean(true),
//   })
//
//   // Add PII filter for credit card numbers with input/output actions
//   guardrail.AddPIIFilter(&PIIFilter{
//   	Type: bedrock.FinancePIIType_CREDIT_DEBIT_CARD_NUMBER(),
//   	Action: bedrock.GuardrailAction_BLOCK,
//   	// below props are optional
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
//   	OutputEnabled: jsii.Boolean(true),
//   })
//
//   // Add PII filter for email addresses
//   guardrail.AddPIIFilter(&PIIFilter{
//   	Type: bedrock.GeneralPIIType_EMAIL(),
//   	Action: bedrock.GuardrailAction_ANONYMIZE,
//   })
//
//   // Add PII filter for US Social Security Numbers
//   guardrail.AddPIIFilter(&PIIFilter{
//   	Type: bedrock.USASpecificPIIType_US_SOCIAL_SECURITY_NUMBER(),
//   	Action: bedrock.GuardrailAction_BLOCK,
//   })
//
//   // Add PII filter for IP addresses
//   guardrail.AddPIIFilter(&PIIFilter{
//   	Type: bedrock.InformationTechnologyPIIType_IP_ADDRESS(),
//   	Action: bedrock.GuardrailAction_ANONYMIZE,
//   })
//
// Experimental.
type GeneralPIIType interface {
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

// The jsii proxy struct for GeneralPIIType
type jsiiProxy_GeneralPIIType struct {
	jsiiProxy_PIIType
}

func (j *jsiiProxy_GeneralPIIType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func GeneralPIIType_ADDRESS() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"ADDRESS",
		&returns,
	)
	return returns
}

func GeneralPIIType_AGE() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"AGE",
		&returns,
	)
	return returns
}

func GeneralPIIType_DRIVER_ID() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"DRIVER_ID",
		&returns,
	)
	return returns
}

func GeneralPIIType_EMAIL() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"EMAIL",
		&returns,
	)
	return returns
}

func GeneralPIIType_LICENSE_PLATE() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"LICENSE_PLATE",
		&returns,
	)
	return returns
}

func GeneralPIIType_NAME() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"NAME",
		&returns,
	)
	return returns
}

func GeneralPIIType_PASSWORD() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"PASSWORD",
		&returns,
	)
	return returns
}

func GeneralPIIType_PHONE() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"PHONE",
		&returns,
	)
	return returns
}

func GeneralPIIType_USERNAME() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"USERNAME",
		&returns,
	)
	return returns
}

func GeneralPIIType_VEHICLE_IDENTIFICATION_NUMBER() GeneralPIIType {
	_init_.Initialize()
	var returns GeneralPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.GeneralPIIType",
		"VEHICLE_IDENTIFICATION_NUMBER",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GeneralPIIType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


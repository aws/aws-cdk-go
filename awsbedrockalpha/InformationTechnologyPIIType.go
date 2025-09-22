package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Types of PII in the domain of IT (Information Technology).
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
type InformationTechnologyPIIType interface {
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

// The jsii proxy struct for InformationTechnologyPIIType
type jsiiProxy_InformationTechnologyPIIType struct {
	jsiiProxy_PIIType
}

func (j *jsiiProxy_InformationTechnologyPIIType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func InformationTechnologyPIIType_AWS_ACCESS_KEY() InformationTechnologyPIIType {
	_init_.Initialize()
	var returns InformationTechnologyPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.InformationTechnologyPIIType",
		"AWS_ACCESS_KEY",
		&returns,
	)
	return returns
}

func InformationTechnologyPIIType_AWS_SECRET_KEY() InformationTechnologyPIIType {
	_init_.Initialize()
	var returns InformationTechnologyPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.InformationTechnologyPIIType",
		"AWS_SECRET_KEY",
		&returns,
	)
	return returns
}

func InformationTechnologyPIIType_IP_ADDRESS() InformationTechnologyPIIType {
	_init_.Initialize()
	var returns InformationTechnologyPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.InformationTechnologyPIIType",
		"IP_ADDRESS",
		&returns,
	)
	return returns
}

func InformationTechnologyPIIType_MAC_ADDRESS() InformationTechnologyPIIType {
	_init_.Initialize()
	var returns InformationTechnologyPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.InformationTechnologyPIIType",
		"MAC_ADDRESS",
		&returns,
	)
	return returns
}

func InformationTechnologyPIIType_URL() InformationTechnologyPIIType {
	_init_.Initialize()
	var returns InformationTechnologyPIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.InformationTechnologyPIIType",
		"URL",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_InformationTechnologyPIIType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


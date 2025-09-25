package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Types of PII in the domain of Finance.
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
type FinancePIIType interface {
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

// The jsii proxy struct for FinancePIIType
type jsiiProxy_FinancePIIType struct {
	jsiiProxy_PIIType
}

func (j *jsiiProxy_FinancePIIType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


func FinancePIIType_CREDIT_DEBIT_CARD_CVV() FinancePIIType {
	_init_.Initialize()
	var returns FinancePIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		"CREDIT_DEBIT_CARD_CVV",
		&returns,
	)
	return returns
}

func FinancePIIType_CREDIT_DEBIT_CARD_EXPIRY() FinancePIIType {
	_init_.Initialize()
	var returns FinancePIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		"CREDIT_DEBIT_CARD_EXPIRY",
		&returns,
	)
	return returns
}

func FinancePIIType_CREDIT_DEBIT_CARD_NUMBER() FinancePIIType {
	_init_.Initialize()
	var returns FinancePIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		"CREDIT_DEBIT_CARD_NUMBER",
		&returns,
	)
	return returns
}

func FinancePIIType_INTERNATIONAL_BANK_ACCOUNT_NUMBER() FinancePIIType {
	_init_.Initialize()
	var returns FinancePIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		"INTERNATIONAL_BANK_ACCOUNT_NUMBER",
		&returns,
	)
	return returns
}

func FinancePIIType_PIN() FinancePIIType {
	_init_.Initialize()
	var returns FinancePIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		"PIN",
		&returns,
	)
	return returns
}

func FinancePIIType_SWIFT_CODE() FinancePIIType {
	_init_.Initialize()
	var returns FinancePIIType
	_jsii_.StaticGet(
		"@aws-cdk/aws-bedrock-alpha.FinancePIIType",
		"SWIFT_CODE",
		&returns,
	)
	return returns
}

func (f *jsiiProxy_FinancePIIType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


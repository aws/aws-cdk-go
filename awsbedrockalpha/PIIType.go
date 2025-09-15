package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Abstract base class for all PII types.
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
type PIIType interface {
	// The string value of the PII type.
	// Experimental.
	Value() *string
	// Returns the string representation of the PII type.
	//
	// Returns: The string value of the PII type.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PIIType
type jsiiProxy_PIIType struct {
	_ byte // padding
}

func (j *jsiiProxy_PIIType) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// The string value of the PII type.
// Experimental.
func NewPIIType_Override(p PIIType, value *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/aws-bedrock-alpha.PIIType",
		[]interface{}{value},
		p,
	)
}

func (p *jsiiProxy_PIIType) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


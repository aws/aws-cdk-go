package awsbedrockalpha


// Interface to define a PII Filter.
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
type PIIFilter struct {
	// The action to take when PII is detected.
	// Experimental.
	Action GuardrailAction `field:"required" json:"action" yaml:"action"`
	// The type of PII to filter.
	// Experimental.
	Type PIIType `field:"required" json:"type" yaml:"type"`
	// The action to take when PII is detected in the input.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the PII filter is enabled for input.
	// Default: true.
	//
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when PII is detected in the output.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the PII filter is enabled for output.
	// Default: true.
	//
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}


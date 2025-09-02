package awsbedrockalpha


// Interface to declare a content filter.
//
// Example:
//   // Create a guardrail to filter inappropriate content
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   	Description: jsii.String("Legal ethical guardrails."),
//   })
//
//   guardrail.AddContentFilter(&ContentFilter{
//   	Type: bedrock.ContentFilterType_SEXUAL,
//   	InputStrength: bedrock.ContentFilterStrength_HIGH,
//   	OutputStrength: bedrock.ContentFilterStrength_MEDIUM,
//   })
//
//   // Create an agent with the guardrail
//   agentWithGuardrail := bedrock.NewAgent(this, jsii.String("AgentWithGuardrail"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_HAIKU_V1_0(),
//   	Instruction: jsii.String("You are a helpful and friendly agent that answers questions about literature."),
//   	Guardrail: guardrail,
//   })
//
// Experimental.
type ContentFilter struct {
	// The strength of the content filter to apply to prompts / user input.
	// Experimental.
	InputStrength ContentFilterStrength `field:"required" json:"inputStrength" yaml:"inputStrength"`
	// The strength of the content filter to apply to model responses.
	// Experimental.
	OutputStrength ContentFilterStrength `field:"required" json:"outputStrength" yaml:"outputStrength"`
	// The type of harmful category that the content filter is applied to.
	// Experimental.
	Type ContentFilterType `field:"required" json:"type" yaml:"type"`
	// The action to take when content is detected in the input.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the content filter is enabled for input.
	// Default: true.
	//
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The input modalities to apply the content filter to.
	// Default: undefined - Applies to text modality.
	//
	// Experimental.
	InputModalities *[]ModalityType `field:"optional" json:"inputModalities" yaml:"inputModalities"`
	// The action to take when content is detected in the output.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the content filter is enabled for output.
	// Default: true.
	//
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
	// The output modalities to apply the content filter to.
	// Default: undefined - Applies to text modality.
	//
	// Experimental.
	OutputModalities *[]ModalityType `field:"optional" json:"outputModalities" yaml:"outputModalities"`
}


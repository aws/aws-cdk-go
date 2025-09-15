package awsbedrockalpha


// The type of modality that can be used in content filters.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   	// Configure tier for content filters (optional)
//   	ContentFiltersTierConfig: bedrock.TierConfig_STANDARD,
//   })
//
//   guardrail.AddContentFilter(&ContentFilter{
//   	Type: bedrock.ContentFilterType_SEXUAL,
//   	InputStrength: bedrock.ContentFilterStrength_HIGH,
//   	OutputStrength: bedrock.ContentFilterStrength_MEDIUM,
//   	// props below are optional
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_NONE,
//   	OutputEnabled: jsii.Boolean(true),
//   	InputModalities: []modalityType{
//   		bedrock.*modalityType_TEXT,
//   		bedrock.*modalityType_IMAGE,
//   	},
//   	OutputModalities: []*modalityType{
//   		bedrock.*modalityType_TEXT,
//   	},
//   })
//
// Experimental.
type ModalityType string

const (
	// Text modality for content filters.
	// Experimental.
	ModalityType_TEXT ModalityType = "TEXT"
	// Image modality for content filters.
	// Experimental.
	ModalityType_IMAGE ModalityType = "IMAGE"
)


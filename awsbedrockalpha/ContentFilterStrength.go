package awsbedrockalpha


// The strength of the content filter.
//
// As you increase the filter strength,
// the likelihood of filtering harmful content increases and the probability
// of seeing harmful content in your application reduces.
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
type ContentFilterStrength string

const (
	// No content filtering applied.
	// Experimental.
	ContentFilterStrength_NONE ContentFilterStrength = "NONE"
	// Low strength content filtering - minimal filtering of harmful content.
	// Experimental.
	ContentFilterStrength_LOW ContentFilterStrength = "LOW"
	// Medium strength content filtering - balanced filtering of harmful content.
	// Experimental.
	ContentFilterStrength_MEDIUM ContentFilterStrength = "MEDIUM"
	// High strength content filtering - aggressive filtering of harmful content.
	// Experimental.
	ContentFilterStrength_HIGH ContentFilterStrength = "HIGH"
)


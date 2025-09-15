package awsbedrockalpha


// The type of harmful category usable in a content filter.
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
type ContentFilterType string

const (
	// Describes input prompts and model responses that indicates sexual interest, activity, or arousal using direct or indirect references to body parts, physical traits, or sex.
	// Experimental.
	ContentFilterType_SEXUAL ContentFilterType = "SEXUAL"
	// Describes input prompts and model responses that includes glorification of or threats to inflict physical pain, hurt, or injury toward a person, group or thing.
	// Experimental.
	ContentFilterType_VIOLENCE ContentFilterType = "VIOLENCE"
	// Describes input prompts and model responses that discriminate, criticize, insult, denounce, or dehumanize a person or group on the basis of an identity (such as race, ethnicity, gender, religion, sexual orientation, ability, and national origin).
	// Experimental.
	ContentFilterType_HATE ContentFilterType = "HATE"
	// Describes input prompts and model responses that includes demeaning, humiliating, mocking, insulting, or belittling language.
	//
	// This type of language is also labeled
	// as bullying.
	// Experimental.
	ContentFilterType_INSULTS ContentFilterType = "INSULTS"
	// Describes input prompts and model responses that seeks or provides information about engaging in misconduct activity, or harming, defrauding, or taking advantage of a person, group or institution.
	// Experimental.
	ContentFilterType_MISCONDUCT ContentFilterType = "MISCONDUCT"
	// Enable to detect and block user inputs attempting to override system instructions.
	//
	// To avoid misclassifying system prompts as a prompt attack and ensure that the filters
	// are selectively applied to user inputs, use input tagging.
	// Experimental.
	ContentFilterType_PROMPT_ATTACK ContentFilterType = "PROMPT_ATTACK"
)


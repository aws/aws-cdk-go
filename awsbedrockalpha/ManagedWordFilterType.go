package awsbedrockalpha


// Managed word list filter types supported by Amazon Bedrock.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   })
//
//   // Add managed word list with input/output actions
//   guardrail.AddManagedWordListFilter(&ManagedWordFilter{
//   	Type: bedrock.ManagedWordFilterType_PROFANITY,
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_NONE,
//   	OutputEnabled: jsii.Boolean(true),
//   })
//
//   // Add individual words
//   guardrail.AddWordFilter(&WordFilter{
//   	Text: jsii.String("drugs"),
//   })
//   guardrail.AddWordFilter(&WordFilter{
//   	Text: jsii.String("competitor"),
//   })
//
//   // Add words from a file
//   guardrail.AddWordFilterFromFile(jsii.String("./scripts/wordsPolicy.csv"))
//
// Experimental.
type ManagedWordFilterType string

const (
	// Filter for profanity and explicit language.
	// Experimental.
	ManagedWordFilterType_PROFANITY ManagedWordFilterType = "PROFANITY"
)


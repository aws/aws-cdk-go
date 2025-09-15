package awsbedrockalpha


// Interface for managed word list filters.
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
type ManagedWordFilter struct {
	// The action to take when a managed word is detected in the input.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the managed word filter is enabled for input.
	// Default: true.
	//
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when a managed word is detected in the output.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the managed word filter is enabled for output.
	// Default: true.
	//
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
	// The type of managed word filter.
	// Default: ManagedWordFilterType.PROFANITY
	//
	// Experimental.
	Type ManagedWordFilterType `field:"optional" json:"type" yaml:"type"`
}


package awsbedrockalpha


// A Regular expression (regex) filter for sensitive information.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   })
//   // Add regex filter with input/output actions
//   guardrail.AddRegexFilter(&RegexFilter{
//   	Name: jsii.String("TestRegexFilter"),
//   	Pattern: jsii.String("test-pattern"),
//   	Action: bedrock.GuardrailAction_ANONYMIZE,
//   	// below props are optional
//   	Description: jsii.String("This is a test regex filter"),
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_ANONYMIZE,
//   	OutputEnabled: jsii.Boolean(true),
//   })
//
// Experimental.
type RegexFilter struct {
	// The action to take when a regex match is detected.
	// Experimental.
	Action GuardrailAction `field:"required" json:"action" yaml:"action"`
	// The name of the regex filter.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The regular expression pattern to match.
	// Experimental.
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The description of the regex filter.
	// Default: - No description.
	//
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The action to take when a regex match is detected in the input.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the regex filter is enabled for input.
	// Default: true.
	//
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when a regex match is detected in the output.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the regex filter is enabled for output.
	// Default: true.
	//
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}


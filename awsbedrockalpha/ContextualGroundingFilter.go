package awsbedrockalpha


// Interface to define a Contextual Grounding Filter.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   })
//   // Add contextual grounding filter with action and enabled flag
//   guardrail.AddContextualGroundingFilter(&ContextualGroundingFilter{
//   	Type: bedrock.ContextualGroundingFilterType_GROUNDING,
//   	Threshold: jsii.Number(0.8),
//   	// the properties below are optional
//   	Action: bedrock.GuardrailAction_BLOCK,
//   	Enabled: jsii.Boolean(true),
//   })
//
// Experimental.
type ContextualGroundingFilter struct {
	// The threshold for the contextual grounding filter.
	//
	// - `0` (blocks nothing)
	// - `0.99` (blocks almost everything)
	// Experimental.
	Threshold *float64 `field:"required" json:"threshold" yaml:"threshold"`
	// The type of contextual grounding filter.
	// Experimental.
	Type ContextualGroundingFilterType `field:"required" json:"type" yaml:"type"`
	// The action to take when contextual grounding is detected.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	Action GuardrailAction `field:"optional" json:"action" yaml:"action"`
	// Whether the contextual grounding filter is enabled.
	// Default: true.
	//
	// Experimental.
	Enabled *bool `field:"optional" json:"enabled" yaml:"enabled"`
}


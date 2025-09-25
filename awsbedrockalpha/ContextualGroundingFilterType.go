package awsbedrockalpha


// The type of contextual grounding filter.
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
type ContextualGroundingFilterType string

const (
	// Grounding score represents the confidence that the model response is factually correct and grounded in the source.
	//
	// If the model response has a lower score than
	// the defined threshold, the response will be blocked and the configured blocked
	// message will be returned to the user. A higher threshold level blocks more responses.
	// Experimental.
	ContextualGroundingFilterType_GROUNDING ContextualGroundingFilterType = "GROUNDING"
	// Relevance score represents the confidence that the model response is relevant to the user's query.
	//
	// If the model response has a lower score than the defined
	// threshold, the response will be blocked and the configured blocked message will
	// be returned to the user. A higher threshold level blocks more responses.
	// Experimental.
	ContextualGroundingFilterType_RELEVANCE ContextualGroundingFilterType = "RELEVANCE"
)


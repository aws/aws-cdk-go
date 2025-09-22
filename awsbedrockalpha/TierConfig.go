package awsbedrockalpha


// ****************************************************************************                            TIER CONFIG ***************************************************************************.
//
// Example:
//   guardrail := bedrock.NewGuardrail(this, jsii.String("bedrockGuardrails"), &GuardrailProps{
//   	GuardrailName: jsii.String("my-BedrockGuardrails"),
//   	// Configure tier for topic filters (optional)
//   	TopicsTierConfig: bedrock.TierConfig_STANDARD,
//   })
//
//   // Use a predefined topic
//   guardrail.AddDeniedTopicFilter(bedrock.Topic_FINANCIAL_ADVICE())
//
//   // Create a custom topic with input/output actions
//   guardrail.AddDeniedTopicFilter(bedrock.Topic_Custom(&CustomTopicProps{
//   	Name: jsii.String("Legal_Advice"),
//   	Definition: jsii.String("Offering guidance or suggestions on legal matters, legal actions, interpretation of laws, or legal rights and responsibilities."),
//   	Examples: []*string{
//   		jsii.String("Can I sue someone for this?"),
//   		jsii.String("What are my legal rights in this situation?"),
//   		jsii.String("Is this action against the law?"),
//   		jsii.String("What should I do to file a legal complaint?"),
//   		jsii.String("Can you explain this law to me?"),
//   	},
//   	// props below are optional
//   	InputAction: bedrock.GuardrailAction_BLOCK,
//   	InputEnabled: jsii.Boolean(true),
//   	OutputAction: bedrock.GuardrailAction_NONE,
//   	OutputEnabled: jsii.Boolean(true),
//   }))
//
// Experimental.
type TierConfig string

const (
	// Provides established guardrails functionality supporting English, French, and Spanish languages.
	// Experimental.
	TierConfig_CLASSIC TierConfig = "CLASSIC"
	// Provides a more robust solution than the CLASSIC tier and has more comprehensive language support.
	//
	// This tier requires that your guardrail use cross-Region inference.
	// Experimental.
	TierConfig_STANDARD TierConfig = "STANDARD"
)


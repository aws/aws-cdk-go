package awsbedrockalpha


// Guardrail action when a sensitive entity is detected.
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
type GuardrailAction string

const (
	// If sensitive information is detected in the prompt or response, the guardrail blocks all the content and returns a message that you configure.
	// Experimental.
	GuardrailAction_BLOCK GuardrailAction = "BLOCK"
	// If sensitive information is detected in the model response, the guardrail masks it with an identifier, the sensitive information is masked and replaced with identifier tags (for example: [NAME-1], [NAME-2], [EMAIL-1], etc.).
	// Experimental.
	GuardrailAction_ANONYMIZE GuardrailAction = "ANONYMIZE"
	// Do not take any action.
	// Experimental.
	GuardrailAction_NONE GuardrailAction = "NONE"
)


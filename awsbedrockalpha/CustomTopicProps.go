package awsbedrockalpha


// Interface for creating a custom Topic.
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
type CustomTopicProps struct {
	// Provide a clear definition to detect and block user inputs and FM responses that fall into this topic.
	//
	// Avoid starting with "don't".
	//
	// Example:
	//   `Investment advice refers to inquiries, guidance, or recommendations
	//   regarding the management or allocation of funds or assets with the goal of
	//   generating returns or achieving specific financial objectives.`
	//
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Representative phrases that refer to the topic.
	//
	// These phrases can represent
	// a user input or a model response. Add between 1 and 100 phrases, up to 100 characters
	// each.
	//
	// Example:
	//   "Where should I invest my money?"
	//
	// Experimental.
	Examples *[]*string `field:"required" json:"examples" yaml:"examples"`
	// The name of the topic to deny.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The action to take when a topic is detected in the input.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	InputAction GuardrailAction `field:"optional" json:"inputAction" yaml:"inputAction"`
	// Whether the topic filter is enabled for input.
	// Default: true.
	//
	// Experimental.
	InputEnabled *bool `field:"optional" json:"inputEnabled" yaml:"inputEnabled"`
	// The action to take when a topic is detected in the output.
	// Default: GuardrailAction.BLOCK
	//
	// Experimental.
	OutputAction GuardrailAction `field:"optional" json:"outputAction" yaml:"outputAction"`
	// Whether the topic filter is enabled for output.
	// Default: true.
	//
	// Experimental.
	OutputEnabled *bool `field:"optional" json:"outputEnabled" yaml:"outputEnabled"`
}


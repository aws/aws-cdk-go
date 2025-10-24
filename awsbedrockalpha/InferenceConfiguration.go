package awsbedrockalpha


// LLM inference configuration.
//
// Example:
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	Instruction: jsii.String("You are a helpful assistant."),
//   	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_FromSteps([]PromptStepConfigBase{
//   		&PromptStepConfigBase{
//   			StepType: bedrock.AgentStepType_PRE_PROCESSING,
//   			StepEnabled: jsii.Boolean(true),
//   			CustomPromptTemplate: jsii.String("Your custom prompt template here"),
//   			InferenceConfig: &InferenceConfiguration{
//   				Temperature: jsii.Number(0),
//   				TopP: jsii.Number(1),
//   				TopK: jsii.Number(250),
//   				MaximumLength: jsii.Number(1),
//   				StopSequences: []*string{
//   					jsii.String("\n\nHuman:"),
//   				},
//   			},
//   		},
//   	}),
//   })
//
// Experimental.
type InferenceConfiguration struct {
	// The maximum number of tokens to generate in the response.
	//
	// Integer
	//
	// min 0
	// max 4096.
	// Experimental.
	MaximumLength *float64 `field:"required" json:"maximumLength" yaml:"maximumLength"`
	// A list of stop sequences.
	//
	// A stop sequence is a sequence of characters that
	// causes the model to stop generating the response.
	//
	// length 0-4.
	// Experimental.
	StopSequences *[]*string `field:"required" json:"stopSequences" yaml:"stopSequences"`
	// The likelihood of the model selecting higher-probability options while generating a response.
	//
	// A lower value makes the model more likely to choose
	// higher-probability options, while a higher value makes the model more
	// likely to choose lower-probability options.
	//
	// Floating point
	//
	// min 0
	// max 1.
	// Experimental.
	Temperature *float64 `field:"required" json:"temperature" yaml:"temperature"`
	// While generating a response, the model determines the probability of the following token at each point of generation.
	//
	// The value that you set for
	// topK is the number of most-likely candidates from which the model chooses
	// the next token in the sequence. For example, if you set topK to 50, the
	// model selects the next token from among the top 50 most likely choices.
	//
	// Integer
	//
	// min 0
	// max 500.
	// Experimental.
	TopK *float64 `field:"required" json:"topK" yaml:"topK"`
	// While generating a response, the model determines the probability of the following token at each point of generation.
	//
	// The value that you set for
	// Top P determines the number of most-likely candidates from which the model
	// chooses the next token in the sequence. For example, if you set topP to
	// 80, the model only selects the next token from the top 80% of the
	// probability distribution of next tokens.
	//
	// Floating point
	//
	// min 0
	// max 1.
	// Experimental.
	TopP *float64 `field:"required" json:"topP" yaml:"topP"`
}


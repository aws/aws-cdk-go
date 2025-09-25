package awsbedrockalpha


// Configuration for the pre-processing step.
//
// Example:
//   parserFunction := lambda.NewFunction(this, jsii.String("ParserFunction"), &FunctionProps{
//   	Runtime: lambda.Runtime_PYTHON_3_10(),
//   	Handler: jsii.String("index.handler"),
//   	Code: lambda.Code_FromAsset(jsii.String("lambda")),
//   })
//
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	Instruction: jsii.String("You are a helpful assistant."),
//   	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_WithCustomParser(&CustomParserProps{
//   		Parser: parserFunction,
//   		PreProcessingStep: &PromptPreProcessingConfigCustomParser{
//   			StepType: bedrock.AgentStepType_PRE_PROCESSING,
//   			UseCustomParser: jsii.Boolean(true),
//   		},
//   	}),
//   })
//
// Experimental.
type PromptPreProcessingConfigCustomParser struct {
	// The type of step this configuration applies to.
	// Experimental.
	StepType AgentStepType `field:"required" json:"stepType" yaml:"stepType"`
	// The custom prompt template to be used.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-placeholders.html
	//
	// Default: - The default prompt template will be used.
	//
	// Experimental.
	CustomPromptTemplate *string `field:"optional" json:"customPromptTemplate" yaml:"customPromptTemplate"`
	// The inference configuration parameters to use.
	// Default: undefined - Default inference configuration will be used.
	//
	// Experimental.
	InferenceConfig *InferenceConfiguration `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
	// Whether to enable or skip this step in the agent sequence.
	// Default: - The default state for each step type is as follows.
	//
	// PRE_PROCESSING – ENABLED
	// ORCHESTRATION – ENABLED
	// KNOWLEDGE_BASE_RESPONSE_GENERATION – ENABLED
	// POST_PROCESSING – DISABLED.
	//
	// Experimental.
	StepEnabled *bool `field:"optional" json:"stepEnabled" yaml:"stepEnabled"`
	// Whether to use the custom Lambda parser defined for the sequence.
	// Default: - false.
	//
	// Experimental.
	UseCustomParser *bool `field:"optional" json:"useCustomParser" yaml:"useCustomParser"`
}


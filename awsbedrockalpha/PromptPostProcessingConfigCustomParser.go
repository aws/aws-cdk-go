package awsbedrockalpha


// Configuration for the post-processing step.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_alpha "github.com/aws/aws-cdk-go/awsbedrockalpha"
//
//   promptPostProcessingConfigCustomParser := &PromptPostProcessingConfigCustomParser{
//   	StepType: bedrock_alpha.AgentStepType_PRE_PROCESSING,
//
//   	// the properties below are optional
//   	CustomPromptTemplate: jsii.String("customPromptTemplate"),
//   	InferenceConfig: &InferenceConfiguration{
//   		MaximumLength: jsii.Number(123),
//   		StopSequences: []*string{
//   			jsii.String("stopSequences"),
//   		},
//   		Temperature: jsii.Number(123),
//   		TopK: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	StepEnabled: jsii.Boolean(false),
//   	UseCustomParser: jsii.Boolean(false),
//   }
//
// Experimental.
type PromptPostProcessingConfigCustomParser struct {
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


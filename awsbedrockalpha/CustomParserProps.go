package awsbedrockalpha

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Properties for configuring a custom Lambda parser for prompt overrides.
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
type CustomParserProps struct {
	// Configuration for the knowledge base response generation step.
	// Default: undefined - No knowledge base response generation configuration.
	//
	// Experimental.
	KnowledgeBaseResponseGenerationStep *PromptKnowledgeBaseResponseGenerationConfigCustomParser `field:"optional" json:"knowledgeBaseResponseGenerationStep" yaml:"knowledgeBaseResponseGenerationStep"`
	// Configuration for the memory summarization step.
	// Default: undefined - No memory summarization configuration.
	//
	// Experimental.
	MemorySummarizationStep *PromptMemorySummarizationConfigCustomParser `field:"optional" json:"memorySummarizationStep" yaml:"memorySummarizationStep"`
	// Configuration for the orchestration step.
	// Default: undefined - No orchestration configuration.
	//
	// Experimental.
	OrchestrationStep *PromptOrchestrationConfigCustomParser `field:"optional" json:"orchestrationStep" yaml:"orchestrationStep"`
	// Lambda function to use as custom parser.
	// Default: undefined - No custom parser is used.
	//
	// Experimental.
	Parser awslambda.IFunction `field:"optional" json:"parser" yaml:"parser"`
	// Configuration for the post-processing step.
	// Default: undefined - No post-processing configuration.
	//
	// Experimental.
	PostProcessingStep *PromptPostProcessingConfigCustomParser `field:"optional" json:"postProcessingStep" yaml:"postProcessingStep"`
	// Configuration for the pre-processing step.
	// Default: undefined - No pre-processing configuration.
	//
	// Experimental.
	PreProcessingStep *PromptPreProcessingConfigCustomParser `field:"optional" json:"preProcessingStep" yaml:"preProcessingStep"`
	// Configuration for the routing classifier step.
	// Default: undefined - No routing classifier configuration.
	//
	// Experimental.
	RoutingClassifierStep *PromptRoutingClassifierConfigCustomParser `field:"optional" json:"routingClassifierStep" yaml:"routingClassifierStep"`
}


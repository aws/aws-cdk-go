package awsbedrockalpha

import (
	_init_ "github.com/aws/aws-cdk-go/awsbedrockalpha/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
)

// Configuration for overriding prompt templates and behaviors in different parts of an agent's sequence.
//
// This allows customizing how the agent processes inputs,
// makes decisions, and generates responses.
//
// Example:
//   agent := bedrock.NewAgent(this, jsii.String("Agent"), &AgentProps{
//   	FoundationModel: bedrock.BedrockFoundationModel_AMAZON_NOVA_LITE_V1(),
//   	Instruction: jsii.String("You are a helpful assistant."),
//   	PromptOverrideConfiguration: bedrock.PromptOverrideConfiguration_FromSteps([]promptStepConfigBase{
//   		&PromptRoutingClassifierConfigCustomParser{
//   			StepType: bedrock.AgentStepType_ROUTING_CLASSIFIER,
//   			StepEnabled: jsii.Boolean(true),
//   			CustomPromptTemplate: jsii.String("Your routing template here"),
//   			FoundationModel: bedrock.BedrockFoundationModel_ANTHROPIC_CLAUDE_V2(),
//   		}.(promptRoutingClassifierConfigCustomParser),
//   	}),
//   })
//
// Experimental.
type PromptOverrideConfiguration interface {
	// Configuration for the knowledge base response generation step.
	// Experimental.
	KnowledgeBaseResponseGenerationStep() *PromptKnowledgeBaseResponseGenerationConfigCustomParser
	// Configuration for the memory summarization step.
	// Experimental.
	MemorySummarizationStep() *PromptMemorySummarizationConfigCustomParser
	// Configuration for the orchestration step.
	// Experimental.
	OrchestrationStep() *PromptOrchestrationConfigCustomParser
	// The custom Lambda parser function to use.
	//
	// The Lambda parser processes and interprets the raw foundation model output.
	// It receives an input event with:
	// - messageVersion: Version of message format (1.0)
	// - agent: Info about the agent (name, id, alias, version)
	// - invokeModelRawResponse: Raw model output to parse
	// - promptType: Type of prompt being parsed
	// - overrideType: Type of override (OUTPUT_PARSER)
	//
	// The Lambda must return a response that the agent uses for next actions.
	// See: https://docs.aws.amazon.com/bedrock/latest/userguide/lambda-parser.html
	//
	// Experimental.
	Parser() awslambda.IFunction
	// Configuration for the post-processing step.
	// Experimental.
	PostProcessingStep() *PromptPostProcessingConfigCustomParser
	// Configuration for the pre-processing step.
	// Experimental.
	PreProcessingStep() *PromptPreProcessingConfigCustomParser
	// Configuration for the routing classifier step.
	// Experimental.
	RoutingClassifierStep() *PromptRoutingClassifierConfigCustomParser
}

// The jsii proxy struct for PromptOverrideConfiguration
type jsiiProxy_PromptOverrideConfiguration struct {
	_ byte // padding
}

func (j *jsiiProxy_PromptOverrideConfiguration) KnowledgeBaseResponseGenerationStep() *PromptKnowledgeBaseResponseGenerationConfigCustomParser {
	var returns *PromptKnowledgeBaseResponseGenerationConfigCustomParser
	_jsii_.Get(
		j,
		"knowledgeBaseResponseGenerationStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) MemorySummarizationStep() *PromptMemorySummarizationConfigCustomParser {
	var returns *PromptMemorySummarizationConfigCustomParser
	_jsii_.Get(
		j,
		"memorySummarizationStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) OrchestrationStep() *PromptOrchestrationConfigCustomParser {
	var returns *PromptOrchestrationConfigCustomParser
	_jsii_.Get(
		j,
		"orchestrationStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) Parser() awslambda.IFunction {
	var returns awslambda.IFunction
	_jsii_.Get(
		j,
		"parser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) PostProcessingStep() *PromptPostProcessingConfigCustomParser {
	var returns *PromptPostProcessingConfigCustomParser
	_jsii_.Get(
		j,
		"postProcessingStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) PreProcessingStep() *PromptPreProcessingConfigCustomParser {
	var returns *PromptPreProcessingConfigCustomParser
	_jsii_.Get(
		j,
		"preProcessingStep",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PromptOverrideConfiguration) RoutingClassifierStep() *PromptRoutingClassifierConfigCustomParser {
	var returns *PromptRoutingClassifierConfigCustomParser
	_jsii_.Get(
		j,
		"routingClassifierStep",
		&returns,
	)
	return returns
}


// Creates a PromptOverrideConfiguration from individual step configurations.
//
// Use this method when you want to override prompts without using a custom parser.
//
// Returns: A new PromptOverrideConfiguration instance.
// Experimental.
func PromptOverrideConfiguration_FromSteps(steps *[]*PromptStepConfigBase) PromptOverrideConfiguration {
	_init_.Initialize()

	if err := validatePromptOverrideConfiguration_FromStepsParameters(steps); err != nil {
		panic(err)
	}
	var returns PromptOverrideConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptOverrideConfiguration",
		"fromSteps",
		[]interface{}{steps},
		&returns,
	)

	return returns
}

// Creates a PromptOverrideConfiguration with a custom Lambda parser function.
// Experimental.
func PromptOverrideConfiguration_WithCustomParser(props *CustomParserProps) PromptOverrideConfiguration {
	_init_.Initialize()

	if err := validatePromptOverrideConfiguration_WithCustomParserParameters(props); err != nil {
		panic(err)
	}
	var returns PromptOverrideConfiguration

	_jsii_.StaticInvoke(
		"@aws-cdk/aws-bedrock-alpha.PromptOverrideConfiguration",
		"withCustomParser",
		[]interface{}{props},
		&returns,
	)

	return returns
}


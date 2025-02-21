package awsbedrock


// Contains configurations to override a prompt template in one part of an agent sequence.
//
// For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptConfigurationProperty := &PromptConfigurationProperty{
//   	BasePromptTemplate: jsii.String("basePromptTemplate"),
//   	InferenceConfiguration: &InferenceConfigurationProperty{
//   		MaximumLength: jsii.Number(123),
//   		StopSequences: []*string{
//   			jsii.String("stopSequences"),
//   		},
//   		Temperature: jsii.Number(123),
//   		TopK: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	ParserMode: jsii.String("parserMode"),
//   	PromptCreationMode: jsii.String("promptCreationMode"),
//   	PromptState: jsii.String("promptState"),
//   	PromptType: jsii.String("promptType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html
//
type CfnAgent_PromptConfigurationProperty struct {
	// Defines the prompt template with which to replace the default prompt template.
	//
	// You can use placeholder variables in the base prompt template to customize the prompt. For more information, see [Prompt template placeholder variables](https://docs.aws.amazon.com/bedrock/latest/userguide/prompt-placeholders.html) . For more information, see [Configure the prompt templates](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts-configure.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-baseprompttemplate
	//
	BasePromptTemplate *string `field:"optional" json:"basePromptTemplate" yaml:"basePromptTemplate"`
	// Contains inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the `promptType` .
	//
	// For more information, see [Inference parameters for foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// Specifies whether to override the default parser Lambda function when parsing the raw foundation model output in the part of the agent sequence defined by the `promptType` .
	//
	// If you set the field as `OVERRIDDEN` , the `overrideLambda` field in the [PromptOverrideConfiguration](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PromptOverrideConfiguration.html) must be specified with the ARN of a Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-parsermode
	//
	ParserMode *string `field:"optional" json:"parserMode" yaml:"parserMode"`
	// Specifies whether to override the default prompt template for this `promptType` .
	//
	// Set this value to `OVERRIDDEN` to use the prompt that you provide in the `basePromptTemplate` . If you leave it as `DEFAULT` , the agent uses a default prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-promptcreationmode
	//
	PromptCreationMode *string `field:"optional" json:"promptCreationMode" yaml:"promptCreationMode"`
	// Specifies whether to allow the inline agent to carry out the step specified in the `promptType` .
	//
	// If you set this value to `DISABLED` , the agent skips that step. The default state for each `promptType` is as follows.
	//
	// - `PRE_PROCESSING` – `ENABLED`
	// - `ORCHESTRATION` – `ENABLED`
	// - `KNOWLEDGE_BASE_RESPONSE_GENERATION` – `ENABLED`
	// - `POST_PROCESSING` – `DISABLED`.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-promptstate
	//
	PromptState *string `field:"optional" json:"promptState" yaml:"promptState"`
	// The step in the agent sequence that this prompt configuration applies to.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptconfiguration.html#cfn-bedrock-agent-promptconfiguration-prompttype
	//
	PromptType *string `field:"optional" json:"promptType" yaml:"promptType"`
}


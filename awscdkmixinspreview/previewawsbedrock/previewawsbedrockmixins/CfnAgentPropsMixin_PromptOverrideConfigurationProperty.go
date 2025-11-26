package previewawsbedrockmixins


// Contains configurations to override prompts in different parts of an agent sequence.
//
// For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   promptOverrideConfigurationProperty := &PromptOverrideConfigurationProperty{
//   	OverrideLambda: jsii.String("overrideLambda"),
//   	PromptConfigurations: []interface{}{
//   		&PromptConfigurationProperty{
//   			AdditionalModelRequestFields: additionalModelRequestFields,
//   			BasePromptTemplate: jsii.String("basePromptTemplate"),
//   			FoundationModel: jsii.String("foundationModel"),
//   			InferenceConfiguration: &InferenceConfigurationProperty{
//   				MaximumLength: jsii.Number(123),
//   				StopSequences: []*string{
//   					jsii.String("stopSequences"),
//   				},
//   				Temperature: jsii.Number(123),
//   				TopK: jsii.Number(123),
//   				TopP: jsii.Number(123),
//   			},
//   			ParserMode: jsii.String("parserMode"),
//   			PromptCreationMode: jsii.String("promptCreationMode"),
//   			PromptState: jsii.String("promptState"),
//   			PromptType: jsii.String("promptType"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptoverrideconfiguration.html
//
type CfnAgentPropsMixin_PromptOverrideConfigurationProperty struct {
	// The ARN of the Lambda function to use when parsing the raw foundation model output in parts of the agent sequence.
	//
	// If you specify this field, at least one of the `promptConfigurations` must contain a `parserMode` value that is set to `OVERRIDDEN` . For more information, see [Parser Lambda function in Amazon Bedrock Agents](https://docs.aws.amazon.com/bedrock/latest/userguide/lambda-parser.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptoverrideconfiguration.html#cfn-bedrock-agent-promptoverrideconfiguration-overridelambda
	//
	OverrideLambda *string `field:"optional" json:"overrideLambda" yaml:"overrideLambda"`
	// Contains configurations to override a prompt template in one part of an agent sequence.
	//
	// For more information, see [Advanced prompts](https://docs.aws.amazon.com/bedrock/latest/userguide/advanced-prompts.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-promptoverrideconfiguration.html#cfn-bedrock-agent-promptoverrideconfiguration-promptconfigurations
	//
	PromptConfigurations interface{} `field:"optional" json:"promptConfigurations" yaml:"promptConfigurations"`
}


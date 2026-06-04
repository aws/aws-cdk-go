package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessModelConfigurationProperty := &HarnessModelConfigurationProperty{
//   	BedrockModelConfig: &HarnessBedrockModelConfigProperty{
//   		ModelId: jsii.String("modelId"),
//
//   		// the properties below are optional
//   		MaxTokens: jsii.Number(123),
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	GeminiModelConfig: &HarnessGeminiModelConfigProperty{
//   		ApiKeyArn: jsii.String("apiKeyArn"),
//   		ModelId: jsii.String("modelId"),
//
//   		// the properties below are optional
//   		MaxTokens: jsii.Number(123),
//   		Temperature: jsii.Number(123),
//   		TopK: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	OpenAiModelConfig: &HarnessOpenAiModelConfigProperty{
//   		ApiKeyArn: jsii.String("apiKeyArn"),
//   		ModelId: jsii.String("modelId"),
//
//   		// the properties below are optional
//   		MaxTokens: jsii.Number(123),
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html
//
type CfnHarness_HarnessModelConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-bedrockmodelconfig
	//
	BedrockModelConfig interface{} `field:"optional" json:"bedrockModelConfig" yaml:"bedrockModelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-geminimodelconfig
	//
	GeminiModelConfig interface{} `field:"optional" json:"geminiModelConfig" yaml:"geminiModelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-openaimodelconfig
	//
	OpenAiModelConfig interface{} `field:"optional" json:"openAiModelConfig" yaml:"openAiModelConfig"`
}


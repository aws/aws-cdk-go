package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessModelConfigurationProperty := &HarnessModelConfigurationProperty{
//   	BedrockModelConfig: &HarnessBedrockModelConfigProperty{
//   		ApiFormat: jsii.String("apiFormat"),
//   		MaxTokens: jsii.Number(123),
//   		ModelId: jsii.String("modelId"),
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	GeminiModelConfig: &HarnessGeminiModelConfigProperty{
//   		ApiKeyArn: jsii.String("apiKeyArn"),
//   		MaxTokens: jsii.Number(123),
//   		ModelId: jsii.String("modelId"),
//   		Temperature: jsii.Number(123),
//   		TopK: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	LiteLlmModelConfig: &HarnessLiteLlmModelConfigProperty{
//   		ApiBase: jsii.String("apiBase"),
//   		ApiKeyArn: jsii.String("apiKeyArn"),
//   		MaxTokens: jsii.Number(123),
//   		ModelId: jsii.String("modelId"),
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   	OpenAiModelConfig: &HarnessOpenAiModelConfigProperty{
//   		ApiFormat: jsii.String("apiFormat"),
//   		ApiKeyArn: jsii.String("apiKeyArn"),
//   		MaxTokens: jsii.Number(123),
//   		ModelId: jsii.String("modelId"),
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html
//
type CfnHarnessPropsMixin_HarnessModelConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-bedrockmodelconfig
	//
	BedrockModelConfig interface{} `field:"optional" json:"bedrockModelConfig" yaml:"bedrockModelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-geminimodelconfig
	//
	GeminiModelConfig interface{} `field:"optional" json:"geminiModelConfig" yaml:"geminiModelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-litellmmodelconfig
	//
	LiteLlmModelConfig interface{} `field:"optional" json:"liteLlmModelConfig" yaml:"liteLlmModelConfig"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessmodelconfiguration.html#cfn-bedrockagentcore-harness-harnessmodelconfiguration-openaimodelconfig
	//
	OpenAiModelConfig interface{} `field:"optional" json:"openAiModelConfig" yaml:"openAiModelConfig"`
}


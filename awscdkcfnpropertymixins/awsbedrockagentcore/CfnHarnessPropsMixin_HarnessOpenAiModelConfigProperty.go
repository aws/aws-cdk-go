package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var additionalParams interface{}
//
//   harnessOpenAiModelConfigProperty := &HarnessOpenAiModelConfigProperty{
//   	AdditionalParams: additionalParams,
//   	ApiBase: jsii.String("apiBase"),
//   	ApiFormat: jsii.String("apiFormat"),
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   	MaxTokens: jsii.Number(123),
//   	ModelId: jsii.String("modelId"),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html
//
type CfnHarnessPropsMixin_HarnessOpenAiModelConfigProperty struct {
	// Provider-specific parameters passed through to the model provider unchanged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-additionalparams
	//
	AdditionalParams interface{} `field:"optional" json:"additionalParams" yaml:"additionalParams"`
	// Optional custom endpoint URL for an OpenAI-compatible endpoint.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-apibase
	//
	ApiBase *string `field:"optional" json:"apiBase" yaml:"apiBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-apiformat
	//
	ApiFormat *string `field:"optional" json:"apiFormat" yaml:"apiFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-apikeyarn
	//
	ApiKeyArn *string `field:"optional" json:"apiKeyArn" yaml:"apiKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


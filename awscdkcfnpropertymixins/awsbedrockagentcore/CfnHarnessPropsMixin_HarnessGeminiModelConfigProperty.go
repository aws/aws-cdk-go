package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var additionalParams interface{}
//
//   harnessGeminiModelConfigProperty := &HarnessGeminiModelConfigProperty{
//   	AdditionalParams: additionalParams,
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   	MaxTokens: jsii.Number(123),
//   	ModelId: jsii.String("modelId"),
//   	Temperature: jsii.Number(123),
//   	TopK: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html
//
type CfnHarnessPropsMixin_HarnessGeminiModelConfigProperty struct {
	// Provider-specific parameters passed through to the Gemini model provider unchanged.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-additionalparams
	//
	AdditionalParams interface{} `field:"optional" json:"additionalParams" yaml:"additionalParams"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-apikeyarn
	//
	ApiKeyArn *string `field:"optional" json:"apiKeyArn" yaml:"apiKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-topk
	//
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessgeminimodelconfig.html#cfn-bedrockagentcore-harness-harnessgeminimodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


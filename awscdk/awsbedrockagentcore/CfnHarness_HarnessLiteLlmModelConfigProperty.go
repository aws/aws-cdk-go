package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessLiteLlmModelConfigProperty := &HarnessLiteLlmModelConfigProperty{
//   	ModelId: jsii.String("modelId"),
//
//   	// the properties below are optional
//   	ApiBase: jsii.String("apiBase"),
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   	MaxTokens: jsii.Number(123),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html
//
type CfnHarness_HarnessLiteLlmModelConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html#cfn-bedrockagentcore-harness-harnesslitellmmodelconfig-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html#cfn-bedrockagentcore-harness-harnesslitellmmodelconfig-apibase
	//
	ApiBase *string `field:"optional" json:"apiBase" yaml:"apiBase"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html#cfn-bedrockagentcore-harness-harnesslitellmmodelconfig-apikeyarn
	//
	ApiKeyArn *string `field:"optional" json:"apiKeyArn" yaml:"apiKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html#cfn-bedrockagentcore-harness-harnesslitellmmodelconfig-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html#cfn-bedrockagentcore-harness-harnesslitellmmodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesslitellmmodelconfig.html#cfn-bedrockagentcore-harness-harnesslitellmmodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


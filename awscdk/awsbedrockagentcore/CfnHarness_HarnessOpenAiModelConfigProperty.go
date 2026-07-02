package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessOpenAiModelConfigProperty := &HarnessOpenAiModelConfigProperty{
//   	ApiKeyArn: jsii.String("apiKeyArn"),
//   	ModelId: jsii.String("modelId"),
//
//   	// the properties below are optional
//   	ApiFormat: jsii.String("apiFormat"),
//   	MaxTokens: jsii.Number(123),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html
//
type CfnHarness_HarnessOpenAiModelConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-apikeyarn
	//
	ApiKeyArn *string `field:"required" json:"apiKeyArn" yaml:"apiKeyArn"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-apiformat
	//
	ApiFormat *string `field:"optional" json:"apiFormat" yaml:"apiFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessopenaimodelconfig.html#cfn-bedrockagentcore-harness-harnessopenaimodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


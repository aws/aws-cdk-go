package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessBedrockModelConfigProperty := &HarnessBedrockModelConfigProperty{
//   	ModelId: jsii.String("modelId"),
//
//   	// the properties below are optional
//   	MaxTokens: jsii.Number(123),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html
//
type CfnHarness_HarnessBedrockModelConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


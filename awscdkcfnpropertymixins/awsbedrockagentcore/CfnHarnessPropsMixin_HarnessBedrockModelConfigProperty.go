package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessBedrockModelConfigProperty := &HarnessBedrockModelConfigProperty{
//   	ApiFormat: jsii.String("apiFormat"),
//   	MaxTokens: jsii.Number(123),
//   	ModelId: jsii.String("modelId"),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html
//
type CfnHarnessPropsMixin_HarnessBedrockModelConfigProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-apiformat
	//
	ApiFormat *string `field:"optional" json:"apiFormat" yaml:"apiFormat"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessbedrockmodelconfig.html#cfn-bedrockagentcore-harness-harnessbedrockmodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


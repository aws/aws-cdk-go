package awsbedrock


// Contains inference configurations related to model inference for a prompt.
//
// For more information, see [Inference parameters](https://docs.aws.amazon.com/bedrock/latest/userguide/inference-parameters.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptModelInferenceConfigurationProperty := &PromptModelInferenceConfigurationProperty{
//   	MaxTokens: jsii.Number(123),
//   	StopSequences: []*string{
//   		jsii.String("stopSequences"),
//   	},
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptmodelinferenceconfiguration.html
//
type CfnFlow_PromptModelInferenceConfigurationProperty struct {
	// The maximum number of tokens to return in the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptmodelinferenceconfiguration.html#cfn-bedrock-flow-promptmodelinferenceconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// A list of strings that define sequences after which the model will stop generating.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptmodelinferenceconfiguration.html#cfn-bedrock-flow-promptmodelinferenceconfiguration-stopsequences
	//
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// Controls the randomness of the response.
	//
	// Choose a lower value for more predictable outputs and a higher value for more surprising outputs.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptmodelinferenceconfiguration.html#cfn-bedrock-flow-promptmodelinferenceconfiguration-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The percentage of most-likely candidates that the model considers for the next token.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptmodelinferenceconfiguration.html#cfn-bedrock-flow-promptmodelinferenceconfiguration-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


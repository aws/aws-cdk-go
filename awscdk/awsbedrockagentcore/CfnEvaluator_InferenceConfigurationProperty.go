package awsbedrockagentcore


// The inference configuration parameters that control model behavior during evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceConfigurationProperty := &InferenceConfigurationProperty{
//   	MaxTokens: jsii.Number(123),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-inferenceconfiguration.html
//
type CfnEvaluator_InferenceConfigurationProperty struct {
	// The maximum number of tokens to generate in the model response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-inferenceconfiguration.html#cfn-bedrockagentcore-evaluator-inferenceconfiguration-maxtokens
	//
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The temperature value that controls randomness in the model's responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-inferenceconfiguration.html#cfn-bedrockagentcore-evaluator-inferenceconfiguration-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The top-p sampling parameter that controls the diversity of the model's responses.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-inferenceconfiguration.html#cfn-bedrockagentcore-evaluator-inferenceconfiguration-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


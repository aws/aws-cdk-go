package awsbedrock


// Contains inference parameters to use when the agent invokes a foundation model in the part of the agent sequence defined by the `promptType` .
//
// For more information, see [Inference parameters for foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html) .
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   inferenceConfigurationProperty := &InferenceConfigurationProperty{
//   	MaximumLength: jsii.Number(123),
//   	StopSequences: []*string{
//   		jsii.String("stopSequences"),
//   	},
//   	Temperature: jsii.Number(123),
//   	TopK: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html
//
type CfnAgent_InferenceConfigurationProperty struct {
	// The maximum number of tokens to allow in the generated response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-maximumlength
	//
	MaximumLength *float64 `field:"optional" json:"maximumLength" yaml:"maximumLength"`
	// A list of stop sequences.
	//
	// A stop sequence is a sequence of characters that causes the model to stop generating the response.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-stopsequences
	//
	StopSequences *[]*string `field:"optional" json:"stopSequences" yaml:"stopSequences"`
	// The likelihood of the model selecting higher-probability options while generating a response.
	//
	// A lower value makes the model more likely to choose higher-probability options, while a higher value makes the model more likely to choose lower-probability options.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// While generating a response, the model determines the probability of the following token at each point of generation.
	//
	// The value that you set for `topK` is the number of most-likely candidates from which the model chooses the next token in the sequence. For example, if you set `topK` to 50, the model selects the next token from among the top 50 most likely choices.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-topk
	//
	TopK *float64 `field:"optional" json:"topK" yaml:"topK"`
	// While generating a response, the model determines the probability of the following token at each point of generation.
	//
	// The value that you set for `Top P` determines the number of most-likely candidates from which the model chooses the next token in the sequence. For example, if you set `topP` to 80, the model only selects the next token from the top 80% of the probability distribution of next tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-agent-inferenceconfiguration.html#cfn-bedrock-agent-inferenceconfiguration-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


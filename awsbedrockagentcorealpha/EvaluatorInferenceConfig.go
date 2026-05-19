package awsbedrockagentcorealpha


// Inference configuration for a custom LLM-as-a-Judge evaluator.
//
// Controls how the foundation model generates evaluation responses.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import bedrock_agentcore_alpha "github.com/aws/aws-cdk-go/awsbedrockagentcorealpha"
//
//   evaluatorInferenceConfig := &EvaluatorInferenceConfig{
//   	MaxTokens: jsii.Number(123),
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// Experimental.
type EvaluatorInferenceConfig struct {
	// The maximum number of tokens to generate in the model response.
	// Default: - The foundation model's default maximum token limit is used.
	//
	// Experimental.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// The temperature value that controls randomness in the model's responses.
	//
	// Higher values produce more diverse outputs. Range: 0.0 to 1.0.
	// Default: - The foundation model's default temperature is used.
	//
	// Experimental.
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The top-p sampling parameter that controls the diversity of the model's responses.
	//
	// Range: 0.0 to 1.0.
	// Default: - The foundation model's default top-p value is used.
	//
	// Experimental.
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


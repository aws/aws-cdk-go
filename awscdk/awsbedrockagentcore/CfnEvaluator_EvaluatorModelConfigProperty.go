package awsbedrockagentcore


// The model configuration that specifies which foundation model to use for evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   evaluatorModelConfigProperty := &EvaluatorModelConfigProperty{
//   	BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   		ModelId: jsii.String("modelId"),
//
//   		// the properties below are optional
//   		AdditionalModelRequestFields: additionalModelRequestFields,
//   		InferenceConfig: &InferenceConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatormodelconfig.html
//
type CfnEvaluator_EvaluatorModelConfigProperty struct {
	// The configuration for using Amazon Bedrock models in evaluator assessments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-evaluatormodelconfig-bedrockevaluatormodelconfig
	//
	BedrockEvaluatorModelConfig interface{} `field:"required" json:"bedrockEvaluatorModelConfig" yaml:"bedrockEvaluatorModelConfig"`
}


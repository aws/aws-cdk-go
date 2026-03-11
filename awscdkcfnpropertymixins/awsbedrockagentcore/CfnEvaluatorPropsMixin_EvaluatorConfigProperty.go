package awsbedrockagentcore


// The configuration that defines how an evaluator assesses agent performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   var additionalModelRequestFields interface{}
//
//   evaluatorConfigProperty := &EvaluatorConfigProperty{
//   	LlmAsAJudge: &LlmAsAJudgeEvaluatorConfigProperty{
//   		Instructions: jsii.String("instructions"),
//   		ModelConfig: &EvaluatorModelConfigProperty{
//   			BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   				AdditionalModelRequestFields: additionalModelRequestFields,
//   				InferenceConfig: &InferenceConfigurationProperty{
//   					MaxTokens: jsii.Number(123),
//   					Temperature: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
//   				ModelId: jsii.String("modelId"),
//   			},
//   		},
//   		RatingScale: &RatingScaleProperty{
//   			Categorical: []interface{}{
//   				&CategoricalScaleDefinitionProperty{
//   					Definition: jsii.String("definition"),
//   					Label: jsii.String("label"),
//   				},
//   			},
//   			Numerical: []interface{}{
//   				&NumericalScaleDefinitionProperty{
//   					Definition: jsii.String("definition"),
//   					Label: jsii.String("label"),
//   					Value: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatorconfig.html
//
type CfnEvaluatorPropsMixin_EvaluatorConfigProperty struct {
	// The configuration for LLM-as-a-Judge evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatorconfig.html#cfn-bedrockagentcore-evaluator-evaluatorconfig-llmasajudge
	//
	LlmAsAJudge interface{} `field:"optional" json:"llmAsAJudge" yaml:"llmAsAJudge"`
}


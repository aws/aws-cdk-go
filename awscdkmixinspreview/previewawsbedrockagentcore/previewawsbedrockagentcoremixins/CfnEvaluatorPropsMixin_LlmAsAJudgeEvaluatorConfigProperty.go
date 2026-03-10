package previewawsbedrockagentcoremixins


// The configuration for LLM-as-a-Judge evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   llmAsAJudgeEvaluatorConfigProperty := &LlmAsAJudgeEvaluatorConfigProperty{
//   	Instructions: jsii.String("instructions"),
//   	ModelConfig: &EvaluatorModelConfigProperty{
//   		BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   			AdditionalModelRequestFields: additionalModelRequestFields,
//   			InferenceConfig: &InferenceConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				Temperature: jsii.Number(123),
//   				TopP: jsii.Number(123),
//   			},
//   			ModelId: jsii.String("modelId"),
//   		},
//   	},
//   	RatingScale: &RatingScaleProperty{
//   		Categorical: []interface{}{
//   			&CategoricalScaleDefinitionProperty{
//   				Definition: jsii.String("definition"),
//   				Label: jsii.String("label"),
//   			},
//   		},
//   		Numerical: []interface{}{
//   			&NumericalScaleDefinitionProperty{
//   				Definition: jsii.String("definition"),
//   				Label: jsii.String("label"),
//   				Value: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig.html
//
type CfnEvaluatorPropsMixin_LlmAsAJudgeEvaluatorConfigProperty struct {
	// The evaluation instructions that guide the language model in assessing agent performance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig.html#cfn-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig-instructions
	//
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// The model configuration that specifies which foundation model to use for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig.html#cfn-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig-modelconfig
	//
	ModelConfig interface{} `field:"optional" json:"modelConfig" yaml:"modelConfig"`
	// The rating scale that defines how evaluators should score agent performance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig.html#cfn-bedrockagentcore-evaluator-llmasajudgeevaluatorconfig-ratingscale
	//
	RatingScale interface{} `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}


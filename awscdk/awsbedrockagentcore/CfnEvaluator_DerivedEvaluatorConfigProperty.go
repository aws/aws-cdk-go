package awsbedrockagentcore


// The configuration for an evaluator derived from an existing base evaluator (a built-in or third-party evaluator), run on your own model.
//
// The base evaluator supplies the prompt and scoring logic.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   derivedEvaluatorConfigProperty := &DerivedEvaluatorConfigProperty{
//   	BaseEvaluatorId: jsii.String("baseEvaluatorId"),
//   	ModelConfig: &EvaluatorModelConfigProperty{
//   		BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   			ModelId: jsii.String("modelId"),
//
//   			// the properties below are optional
//   			AdditionalModelRequestFields: additionalModelRequestFields,
//   			InferenceConfig: &InferenceConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				Temperature: jsii.Number(123),
//   				TopP: jsii.Number(123),
//   			},
//   		},
//   		ResponsesEvaluatorModelConfig: &OpenResponsesEvaluatorModelConfigProperty{
//   			ModelId: jsii.String("modelId"),
//
//   			// the properties below are optional
//   			MaxOutputTokens: jsii.Number(123),
//   			Reasoning: &ReasoningConfigurationProperty{
//   				Effort: jsii.String("effort"),
//   			},
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-derivedevaluatorconfig.html
//
type CfnEvaluator_DerivedEvaluatorConfigProperty struct {
	// The identifier of the base evaluator whose logic to run (a Builtin.* or ThirdParty.* evaluator).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-derivedevaluatorconfig.html#cfn-bedrockagentcore-evaluator-derivedevaluatorconfig-baseevaluatorid
	//
	BaseEvaluatorId *string `field:"required" json:"baseEvaluatorId" yaml:"baseEvaluatorId"`
	// The model configuration that specifies which foundation model to use for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-derivedevaluatorconfig.html#cfn-bedrockagentcore-evaluator-derivedevaluatorconfig-modelconfig
	//
	ModelConfig interface{} `field:"required" json:"modelConfig" yaml:"modelConfig"`
}


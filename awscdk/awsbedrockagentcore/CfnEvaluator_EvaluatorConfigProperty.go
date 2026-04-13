package awsbedrockagentcore


// The configuration that defines how an evaluator assesses agent performance.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   evaluatorConfigProperty := &EvaluatorConfigProperty{
//   	CodeBased: &CodeBasedEvaluatorConfigProperty{
//   		LambdaConfig: &LambdaEvaluatorConfigProperty{
//   			LambdaArn: jsii.String("lambdaArn"),
//
//   			// the properties below are optional
//   			LambdaTimeoutInSeconds: jsii.Number(123),
//   		},
//   	},
//   	LlmAsAJudge: &LlmAsAJudgeEvaluatorConfigProperty{
//   		Instructions: jsii.String("instructions"),
//   		ModelConfig: &EvaluatorModelConfigProperty{
//   			BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   				ModelId: jsii.String("modelId"),
//
//   				// the properties below are optional
//   				AdditionalModelRequestFields: additionalModelRequestFields,
//   				InferenceConfig: &InferenceConfigurationProperty{
//   					MaxTokens: jsii.Number(123),
//   					Temperature: jsii.Number(123),
//   					TopP: jsii.Number(123),
//   				},
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
type CfnEvaluator_EvaluatorConfigProperty struct {
	// The configuration for code-based evaluation using a Lambda function.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatorconfig.html#cfn-bedrockagentcore-evaluator-evaluatorconfig-codebased
	//
	CodeBased interface{} `field:"optional" json:"codeBased" yaml:"codeBased"`
	// The configuration for LLM-as-a-Judge evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatorconfig.html#cfn-bedrockagentcore-evaluator-evaluatorconfig-llmasajudge
	//
	LlmAsAJudge interface{} `field:"optional" json:"llmAsAJudge" yaml:"llmAsAJudge"`
}


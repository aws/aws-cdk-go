package awsbedrockagentcore


// The configuration for using OpenResponses-compatible models in evaluator assessments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   openResponsesEvaluatorModelConfigProperty := &OpenResponsesEvaluatorModelConfigProperty{
//   	ModelId: jsii.String("modelId"),
//
//   	// the properties below are optional
//   	MaxOutputTokens: jsii.Number(123),
//   	Reasoning: &ReasoningConfigurationProperty{
//   		Effort: jsii.String("effort"),
//   	},
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html
//
type CfnEvaluator_OpenResponsesEvaluatorModelConfigProperty struct {
	// The identifier of the model to use for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// The maximum number of output tokens to generate, including visible output and reasoning tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-maxoutputtokens
	//
	MaxOutputTokens *float64 `field:"optional" json:"maxOutputTokens" yaml:"maxOutputTokens"`
	// The reasoning configuration for reasoning models.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-reasoning
	//
	Reasoning interface{} `field:"optional" json:"reasoning" yaml:"reasoning"`
	// The sampling temperature between 0 and 2.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-temperature
	//
	Temperature *float64 `field:"optional" json:"temperature" yaml:"temperature"`
	// The nucleus sampling probability mass between 0 and 1.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-topp
	//
	TopP *float64 `field:"optional" json:"topP" yaml:"topP"`
}


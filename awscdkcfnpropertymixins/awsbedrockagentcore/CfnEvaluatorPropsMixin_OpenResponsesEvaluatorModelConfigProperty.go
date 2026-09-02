package awsbedrockagentcore


// The configuration for using OpenResponses-compatible models in evaluator assessments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   openResponsesEvaluatorModelConfigProperty := &OpenResponsesEvaluatorModelConfigProperty{
//   	MaxOutputTokens: jsii.Number(123),
//   	ModelId: jsii.String("modelId"),
//   	Reasoning: &ReasoningConfigurationProperty{
//   		Effort: jsii.String("effort"),
//   	},
//   	Temperature: jsii.Number(123),
//   	TopP: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html
//
type CfnEvaluatorPropsMixin_OpenResponsesEvaluatorModelConfigProperty struct {
	// The maximum number of output tokens to generate, including visible output and reasoning tokens.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-maxoutputtokens
	//
	MaxOutputTokens *float64 `field:"optional" json:"maxOutputTokens" yaml:"maxOutputTokens"`
	// The identifier of the model to use for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-openresponsesevaluatormodelconfig-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
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


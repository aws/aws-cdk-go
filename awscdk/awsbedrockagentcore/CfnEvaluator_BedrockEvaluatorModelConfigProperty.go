package awsbedrockagentcore


// The configuration for using Amazon Bedrock models in evaluator assessments.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
//
//   bedrockEvaluatorModelConfigProperty := &BedrockEvaluatorModelConfigProperty{
//   	ModelId: jsii.String("modelId"),
//
//   	// the properties below are optional
//   	AdditionalModelRequestFields: additionalModelRequestFields,
//   	InferenceConfig: &InferenceConfigurationProperty{
//   		MaxTokens: jsii.Number(123),
//   		Temperature: jsii.Number(123),
//   		TopP: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-bedrockevaluatormodelconfig.html
//
type CfnEvaluator_BedrockEvaluatorModelConfigProperty struct {
	// The identifier of the Amazon Bedrock model to use for evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-bedrockevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-bedrockevaluatormodelconfig-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Additional model-specific request fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-bedrockevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-bedrockevaluatormodelconfig-additionalmodelrequestfields
	//
	AdditionalModelRequestFields interface{} `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// The inference configuration parameters that control model behavior during evaluation.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-bedrockevaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-bedrockevaluatormodelconfig-inferenceconfig
	//
	InferenceConfig interface{} `field:"optional" json:"inferenceConfig" yaml:"inferenceConfig"`
}


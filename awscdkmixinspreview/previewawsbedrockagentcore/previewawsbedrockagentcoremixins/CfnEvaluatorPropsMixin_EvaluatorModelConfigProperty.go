package previewawsbedrockagentcoremixins


// The model configuration that specifies which foundation model to use for evaluation.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//
//   evaluatorModelConfigProperty := &EvaluatorModelConfigProperty{
//   	BedrockEvaluatorModelConfig: &BedrockEvaluatorModelConfigProperty{
//   		AdditionalModelRequestFields: additionalModelRequestFields,
//   		InferenceConfig: &InferenceConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   		ModelId: jsii.String("modelId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatormodelconfig.html
//
type CfnEvaluatorPropsMixin_EvaluatorModelConfigProperty struct {
	// The configuration for using Amazon Bedrock models in evaluator assessments.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-evaluator-evaluatormodelconfig.html#cfn-bedrockagentcore-evaluator-evaluatormodelconfig-bedrockevaluatormodelconfig
	//
	BedrockEvaluatorModelConfig interface{} `field:"optional" json:"bedrockEvaluatorModelConfig" yaml:"bedrockEvaluatorModelConfig"`
}


package previewawsbedrockmixins


// Contains configurations for a prompt defined inline in the node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptFlowNodeInlineConfigurationProperty := &PromptFlowNodeInlineConfigurationProperty{
//   	InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   		Text: &PromptModelInferenceConfigurationProperty{
//   			MaxTokens: jsii.Number(123),
//   			StopSequences: []*string{
//   				jsii.String("stopSequences"),
//   			},
//   			Temperature: jsii.Number(123),
//   			TopP: jsii.Number(123),
//   		},
//   	},
//   	ModelId: jsii.String("modelId"),
//   	TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   		Text: &TextPromptTemplateConfigurationProperty{
//   			InputVariables: []interface{}{
//   				&PromptInputVariableProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Text: jsii.String("text"),
//   		},
//   	},
//   	TemplateType: jsii.String("templateType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html
//
type CfnFlowVersionPropsMixin_PromptFlowNodeInlineConfigurationProperty struct {
	// Contains inference configurations for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) to run inference with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// Contains a prompt and variables in the prompt that can be replaced with values at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// The type of prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-templatetype
	//
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}


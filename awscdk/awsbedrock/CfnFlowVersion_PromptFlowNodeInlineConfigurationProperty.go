package awsbedrock


// Contains configurations for a prompt defined inline in the node.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptFlowNodeInlineConfigurationProperty := &PromptFlowNodeInlineConfigurationProperty{
//   	ModelId: jsii.String("modelId"),
//   	TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   		Text: &TextPromptTemplateConfigurationProperty{
//   			Text: jsii.String("text"),
//
//   			// the properties below are optional
//   			InputVariables: []interface{}{
//   				&PromptInputVariableProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   		},
//   	},
//   	TemplateType: jsii.String("templateType"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html
//
type CfnFlowVersion_PromptFlowNodeInlineConfigurationProperty struct {
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) to run inference with.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-modelid
	//
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Contains a prompt and variables in the prompt that can be replaced with values at runtime.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"required" json:"templateConfiguration" yaml:"templateConfiguration"`
	// The type of prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-templatetype
	//
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Contains inference configurations for the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodeinlineconfiguration.html#cfn-bedrock-flowversion-promptflownodeinlineconfiguration-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
}


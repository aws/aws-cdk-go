package mixinsawsbedrock


// Contains configurations for a prompt and whether it is from Prompt management or defined inline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   promptFlowNodeSourceConfigurationProperty := &PromptFlowNodeSourceConfigurationProperty{
//   	Inline: &PromptFlowNodeInlineConfigurationProperty{
//   		InferenceConfiguration: &PromptInferenceConfigurationProperty{
//   			Text: &PromptModelInferenceConfigurationProperty{
//   				MaxTokens: jsii.Number(123),
//   				StopSequences: []*string{
//   					jsii.String("stopSequences"),
//   				},
//   				Temperature: jsii.Number(123),
//   				TopP: jsii.Number(123),
//   			},
//   		},
//   		ModelId: jsii.String("modelId"),
//   		TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   			Text: &TextPromptTemplateConfigurationProperty{
//   				InputVariables: []interface{}{
//   					&PromptInputVariableProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Text: jsii.String("text"),
//   			},
//   		},
//   		TemplateType: jsii.String("templateType"),
//   	},
//   	Resource: &PromptFlowNodeResourceConfigurationProperty{
//   		PromptArn: jsii.String("promptArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptflownodesourceconfiguration.html
//
type CfnFlowPropsMixin_PromptFlowNodeSourceConfigurationProperty struct {
	// Contains configurations for a prompt that is defined inline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptflownodesourceconfiguration.html#cfn-bedrock-flow-promptflownodesourceconfiguration-inline
	//
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// Contains configurations for a prompt from Prompt management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-promptflownodesourceconfiguration.html#cfn-bedrock-flow-promptflownodesourceconfiguration-resource
	//
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}


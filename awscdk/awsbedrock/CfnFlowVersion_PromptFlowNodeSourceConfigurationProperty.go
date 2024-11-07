package awsbedrock


// Contains configurations for a prompt and whether it is from Prompt management or defined inline.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   promptFlowNodeSourceConfigurationProperty := &PromptFlowNodeSourceConfigurationProperty{
//   	Inline: &PromptFlowNodeInlineConfigurationProperty{
//   		ModelId: jsii.String("modelId"),
//   		TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   			Text: &TextPromptTemplateConfigurationProperty{
//   				Text: jsii.String("text"),
//
//   				// the properties below are optional
//   				InputVariables: []interface{}{
//   					&PromptInputVariableProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   		TemplateType: jsii.String("templateType"),
//
//   		// the properties below are optional
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
//   	},
//   	Resource: &PromptFlowNodeResourceConfigurationProperty{
//   		PromptArn: jsii.String("promptArn"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodesourceconfiguration.html
//
type CfnFlowVersion_PromptFlowNodeSourceConfigurationProperty struct {
	// Contains configurations for a prompt that is defined inline.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodesourceconfiguration.html#cfn-bedrock-flowversion-promptflownodesourceconfiguration-inline
	//
	Inline interface{} `field:"optional" json:"inline" yaml:"inline"`
	// Contains configurations for a prompt from Prompt management.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flowversion-promptflownodesourceconfiguration.html#cfn-bedrock-flowversion-promptflownodesourceconfiguration-resource
	//
	Resource interface{} `field:"optional" json:"resource" yaml:"resource"`
}


package awsbedrock


// Contains details about a variant of the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   promptVariantProperty := &PromptVariantProperty{
//   	Name: jsii.String("name"),
//   	TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   		Chat: &ChatPromptTemplateConfigurationProperty{
//   			Messages: []interface{}{
//   				&MessageProperty{
//   					Content: []interface{}{
//   						&ContentBlockProperty{
//   							Text: jsii.String("text"),
//   						},
//   					},
//   					Role: jsii.String("role"),
//   				},
//   			},
//
//   			// the properties below are optional
//   			InputVariables: []interface{}{
//   				&PromptInputVariableProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			System: []interface{}{
//   				&SystemContentBlockProperty{
//   					Text: jsii.String("text"),
//   				},
//   			},
//   			ToolConfiguration: &ToolConfigurationProperty{
//   				Tools: []interface{}{
//   					&ToolProperty{
//   						ToolSpec: &ToolSpecificationProperty{
//   							InputSchema: &ToolInputSchemaProperty{
//   								Json: json,
//   							},
//   							Name: jsii.String("name"),
//
//   							// the properties below are optional
//   							Description: jsii.String("description"),
//   						},
//   					},
//   				},
//
//   				// the properties below are optional
//   				ToolChoice: &ToolChoiceProperty{
//   					Any: any,
//   					Auto: auto,
//   					Tool: &SpecificToolChoiceProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
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
//   	GenAiResource: &PromptGenAiResourceProperty{
//   		Agent: &PromptAgentResourceProperty{
//   			AgentIdentifier: jsii.String("agentIdentifier"),
//   		},
//   	},
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html
//
type CfnPromptVersion_PromptVariantProperty struct {
	// The name of the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Contains configurations for the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"required" json:"templateConfiguration" yaml:"templateConfiguration"`
	// The type of prompt template to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-templatetype
	//
	TemplateType *string `field:"required" json:"templateType" yaml:"templateType"`
	// Specifies a generative AI resource with which to use the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-genairesource
	//
	GenAiResource interface{} `field:"optional" json:"genAiResource" yaml:"genAiResource"`
	// Contains inference configurations for the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) with which to run inference on the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}


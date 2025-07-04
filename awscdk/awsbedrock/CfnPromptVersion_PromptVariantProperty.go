package awsbedrock


// Contains details about a variant of the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var additionalModelRequestFields interface{}
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
//   							CachePoint: &CachePointBlockProperty{
//   								Type: jsii.String("type"),
//   							},
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
//   					CachePoint: &CachePointBlockProperty{
//   						Type: jsii.String("type"),
//   					},
//   					Text: jsii.String("text"),
//   				},
//   			},
//   			ToolConfiguration: &ToolConfigurationProperty{
//   				Tools: []interface{}{
//   					&ToolProperty{
//   						CachePoint: &CachePointBlockProperty{
//   							Type: jsii.String("type"),
//   						},
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
//   			CachePoint: &CachePointBlockProperty{
//   				Type: jsii.String("type"),
//   			},
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
//   	AdditionalModelRequestFields: additionalModelRequestFields,
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
//   	Metadata: []interface{}{
//   		&PromptMetadataEntryProperty{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
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
	// Contains model-specific inference configurations that aren't in the `inferenceConfiguration` field.
	//
	// To see model-specific inference parameters, see [Inference request parameters and response fields for foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-additionalmodelrequestfields
	//
	AdditionalModelRequestFields interface{} `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// Specifies a generative AI resource with which to use the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-genairesource
	//
	GenAiResource interface{} `field:"optional" json:"genAiResource" yaml:"genAiResource"`
	// Contains inference configurations for the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// An array of objects, each containing a key-value pair that defines a metadata tag and value to attach to a prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) with which to run inference on the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-promptversion-promptvariant.html#cfn-bedrock-promptversion-promptvariant-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
}


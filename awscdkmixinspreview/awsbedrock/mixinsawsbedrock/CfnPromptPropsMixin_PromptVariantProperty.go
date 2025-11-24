package mixinsawsbedrock


// Contains details about a variant of the prompt.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var additionalModelRequestFields interface{}
//   var any interface{}
//   var auto interface{}
//   var json interface{}
//
//   promptVariantProperty := &PromptVariantProperty{
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
//   	Name: jsii.String("name"),
//   	TemplateConfiguration: &PromptTemplateConfigurationProperty{
//   		Chat: &ChatPromptTemplateConfigurationProperty{
//   			InputVariables: []interface{}{
//   				&PromptInputVariableProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
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
//   			System: []interface{}{
//   				&SystemContentBlockProperty{
//   					CachePoint: &CachePointBlockProperty{
//   						Type: jsii.String("type"),
//   					},
//   					Text: jsii.String("text"),
//   				},
//   			},
//   			ToolConfiguration: &ToolConfigurationProperty{
//   				ToolChoice: &ToolChoiceProperty{
//   					Any: any,
//   					Auto: auto,
//   					Tool: &SpecificToolChoiceProperty{
//   						Name: jsii.String("name"),
//   					},
//   				},
//   				Tools: []interface{}{
//   					&ToolProperty{
//   						CachePoint: &CachePointBlockProperty{
//   							Type: jsii.String("type"),
//   						},
//   						ToolSpec: &ToolSpecificationProperty{
//   							Description: jsii.String("description"),
//   							InputSchema: &ToolInputSchemaProperty{
//   								Json: json,
//   							},
//   							Name: jsii.String("name"),
//   						},
//   					},
//   				},
//   			},
//   		},
//   		Text: &TextPromptTemplateConfigurationProperty{
//   			CachePoint: &CachePointBlockProperty{
//   				Type: jsii.String("type"),
//   			},
//   			InputVariables: []interface{}{
//   				&PromptInputVariableProperty{
//   					Name: jsii.String("name"),
//   				},
//   			},
//   			Text: jsii.String("text"),
//   			TextS3Location: &TextS3LocationProperty{
//   				Bucket: jsii.String("bucket"),
//   				Key: jsii.String("key"),
//   				Version: jsii.String("version"),
//   			},
//   		},
//   	},
//   	TemplateType: jsii.String("templateType"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html
//
type CfnPromptPropsMixin_PromptVariantProperty struct {
	// Contains model-specific inference configurations that aren't in the `inferenceConfiguration` field.
	//
	// To see model-specific inference parameters, see [Inference request parameters and response fields for foundation models](https://docs.aws.amazon.com/bedrock/latest/userguide/model-parameters.html) .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-additionalmodelrequestfields
	//
	AdditionalModelRequestFields interface{} `field:"optional" json:"additionalModelRequestFields" yaml:"additionalModelRequestFields"`
	// Specifies a generative AI resource with which to use the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-genairesource
	//
	GenAiResource interface{} `field:"optional" json:"genAiResource" yaml:"genAiResource"`
	// Contains inference configurations for the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-inferenceconfiguration
	//
	InferenceConfiguration interface{} `field:"optional" json:"inferenceConfiguration" yaml:"inferenceConfiguration"`
	// An array of objects, each containing a key-value pair that defines a metadata tag and value to attach to a prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-metadata
	//
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// The unique identifier of the model or [inference profile](https://docs.aws.amazon.com/bedrock/latest/userguide/cross-region-inference.html) with which to run inference on the prompt.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// The name of the prompt variant.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contains configurations for the prompt template.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-templateconfiguration
	//
	TemplateConfiguration interface{} `field:"optional" json:"templateConfiguration" yaml:"templateConfiguration"`
	// The type of prompt template to use.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-prompt-promptvariant.html#cfn-bedrock-prompt-promptvariant-templatetype
	//
	TemplateType *string `field:"optional" json:"templateType" yaml:"templateType"`
}


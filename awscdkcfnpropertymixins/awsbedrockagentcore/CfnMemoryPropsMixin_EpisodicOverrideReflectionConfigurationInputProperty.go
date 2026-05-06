package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   episodicOverrideReflectionConfigurationInputProperty := &EpisodicOverrideReflectionConfigurationInputProperty{
//   	AppendToPrompt: jsii.String("appendToPrompt"),
//   	MemoryRecordSchema: &MemoryRecordSchemaProperty{
//   		MetadataSchema: []interface{}{
//   			&MetadataSchemaEntryProperty{
//   				ExtractionConfig: &ExtractionConfigProperty{
//   					LlmExtractionConfig: &LlmExtractionConfigProperty{
//   						Definition: jsii.String("definition"),
//   						LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   						Validation: &ValidationProperty{
//   							NumberValidation: &NumberValidationProperty{
//   								MaxValue: jsii.Number(123),
//   								MinValue: jsii.Number(123),
//   							},
//   							StringListValidation: &StringListValidationProperty{
//   								AllowedValues: []*string{
//   									jsii.String("allowedValues"),
//   								},
//   								MaxItems: jsii.Number(123),
//   							},
//   							StringValidation: &StringValidationProperty{
//   								AllowedValues: []*string{
//   									jsii.String("allowedValues"),
//   								},
//   							},
//   						},
//   					},
//   				},
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	ModelId: jsii.String("modelId"),
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   	NamespaceTemplates: []*string{
//   		jsii.String("namespaceTemplates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html
//
type CfnMemoryPropsMixin_EpisodicOverrideReflectionConfigurationInputProperty struct {
	// Text prompt for model instructions.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-appendtoprompt
	//
	AppendToPrompt *string `field:"optional" json:"appendToPrompt" yaml:"appendToPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-memoryrecordschema
	//
	MemoryRecordSchema interface{} `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-modelid
	//
	ModelId *string `field:"optional" json:"modelId" yaml:"modelId"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicoverridereflectionconfigurationinput-namespacetemplates
	//
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
}


package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   episodicReflectionConfigurationInputProperty := &EpisodicReflectionConfigurationInputProperty{
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
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   	NamespaceTemplates: []*string{
//   		jsii.String("namespaceTemplates"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html
//
type CfnMemoryPropsMixin_EpisodicReflectionConfigurationInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicreflectionconfigurationinput-memoryrecordschema
	//
	MemoryRecordSchema interface{} `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicreflectionconfigurationinput-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicreflectionconfigurationinput.html#cfn-bedrockagentcore-memory-episodicreflectionconfigurationinput-namespacetemplates
	//
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
}


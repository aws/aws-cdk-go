package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   episodicOverrideProperty := &EpisodicOverrideProperty{
//   	Consolidation: &EpisodicOverrideConsolidationConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   	Extraction: &EpisodicOverrideExtractionConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   	Reflection: &EpisodicOverrideReflectionConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//
//   		// the properties below are optional
//   		MemoryRecordSchema: &MemoryRecordSchemaProperty{
//   			MetadataSchema: []interface{}{
//   				&MetadataSchemaEntryProperty{
//   					Key: jsii.String("key"),
//
//   					// the properties below are optional
//   					ExtractionConfig: &ExtractionConfigProperty{
//   						LlmExtractionConfig: &LlmExtractionConfigProperty{
//   							Definition: jsii.String("definition"),
//
//   							// the properties below are optional
//   							LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   							Validation: &ValidationProperty{
//   								NumberValidation: &NumberValidationProperty{
//   									MaxValue: jsii.Number(123),
//   									MinValue: jsii.Number(123),
//   								},
//   								StringListValidation: &StringListValidationProperty{
//   									AllowedValues: []*string{
//   										jsii.String("allowedValues"),
//   									},
//   									MaxItems: jsii.Number(123),
//   								},
//   								StringValidation: &StringValidationProperty{
//   									AllowedValues: []*string{
//   										jsii.String("allowedValues"),
//   									},
//   								},
//   							},
//   						},
//   					},
//   					Type: jsii.String("type"),
//   				},
//   			},
//   		},
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		NamespaceTemplates: []*string{
//   			jsii.String("namespaceTemplates"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html
//
type CfnMemory_EpisodicOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html#cfn-bedrockagentcore-memory-episodicoverride-consolidation
	//
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html#cfn-bedrockagentcore-memory-episodicoverride-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicoverride.html#cfn-bedrockagentcore-memory-episodicoverride-reflection
	//
	Reflection interface{} `field:"optional" json:"reflection" yaml:"reflection"`
}


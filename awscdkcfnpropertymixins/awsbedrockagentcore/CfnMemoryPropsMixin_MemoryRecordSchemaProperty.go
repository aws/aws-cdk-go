package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   memoryRecordSchemaProperty := &MemoryRecordSchemaProperty{
//   	MetadataSchema: []interface{}{
//   		&MetadataSchemaEntryProperty{
//   			ExtractionConfig: &ExtractionConfigProperty{
//   				LlmExtractionConfig: &LlmExtractionConfigProperty{
//   					Definition: jsii.String("definition"),
//   					LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   					Validation: &ValidationProperty{
//   						NumberValidation: &NumberValidationProperty{
//   							MaxValue: jsii.Number(123),
//   							MinValue: jsii.Number(123),
//   						},
//   						StringListValidation: &StringListValidationProperty{
//   							AllowedValues: []*string{
//   								jsii.String("allowedValues"),
//   							},
//   							MaxItems: jsii.Number(123),
//   						},
//   						StringValidation: &StringValidationProperty{
//   							AllowedValues: []*string{
//   								jsii.String("allowedValues"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memoryrecordschema.html
//
type CfnMemoryPropsMixin_MemoryRecordSchemaProperty struct {
	// List of metadata schema entries.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memoryrecordschema.html#cfn-bedrockagentcore-memory-memoryrecordschema-metadataschema
	//
	MetadataSchema interface{} `field:"optional" json:"metadataSchema" yaml:"metadataSchema"`
}


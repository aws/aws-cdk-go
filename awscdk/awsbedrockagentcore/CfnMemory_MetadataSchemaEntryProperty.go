package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   metadataSchemaEntryProperty := &MetadataSchemaEntryProperty{
//   	Key: jsii.String("key"),
//
//   	// the properties below are optional
//   	ExtractionConfig: &ExtractionConfigProperty{
//   		LlmExtractionConfig: &LlmExtractionConfigProperty{
//   			Definition: jsii.String("definition"),
//
//   			// the properties below are optional
//   			LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   			Validation: &ValidationProperty{
//   				NumberValidation: &NumberValidationProperty{
//   					MaxValue: jsii.Number(123),
//   					MinValue: jsii.Number(123),
//   				},
//   				StringListValidation: &StringListValidationProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   					MaxItems: jsii.Number(123),
//   				},
//   				StringValidation: &StringValidationProperty{
//   					AllowedValues: []*string{
//   						jsii.String("allowedValues"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	Type: jsii.String("type"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-metadataschemaentry.html
//
type CfnMemory_MetadataSchemaEntryProperty struct {
	// Key name for metadata fields.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-metadataschemaentry.html#cfn-bedrockagentcore-memory-metadataschemaentry-key
	//
	Key *string `field:"required" json:"key" yaml:"key"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-metadataschemaentry.html#cfn-bedrockagentcore-memory-metadataschemaentry-extractionconfig
	//
	ExtractionConfig interface{} `field:"optional" json:"extractionConfig" yaml:"extractionConfig"`
	// Supported data types for metadata values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-metadataschemaentry.html#cfn-bedrockagentcore-memory-metadataschemaentry-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}


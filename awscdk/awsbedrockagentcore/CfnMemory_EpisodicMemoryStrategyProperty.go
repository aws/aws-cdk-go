package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   episodicMemoryStrategyProperty := &EpisodicMemoryStrategyProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
//   	MemoryRecordSchema: &MemoryRecordSchemaProperty{
//   		MetadataSchema: []interface{}{
//   			&MetadataSchemaEntryProperty{
//   				Key: jsii.String("key"),
//
//   				// the properties below are optional
//   				ExtractionConfig: &ExtractionConfigProperty{
//   					LlmExtractionConfig: &LlmExtractionConfigProperty{
//   						Definition: jsii.String("definition"),
//
//   						// the properties below are optional
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
//   				ExtractionType: jsii.String("extractionType"),
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
//   	RecordExpiryDurationDays: jsii.Number(123),
//   	RecordExpirySpecification: &RecordExpirySpecificationProperty{
//   		RecordExpiryDurationDays: jsii.Number(123),
//   	},
//   	ReflectionConfiguration: &EpisodicReflectionConfigurationInputProperty{
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
//   					ExtractionType: jsii.String("extractionType"),
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
//   	Status: jsii.String("status"),
//   	StrategyId: jsii.String("strategyId"),
//   	Type: jsii.String("type"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html
//
type CfnMemory_EpisodicMemoryStrategyProperty struct {
	// Name of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Description of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-memoryrecordschema
	//
	MemoryRecordSchema interface{} `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-namespacetemplates
	//
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
	// The duration, in days, after which long-term memory records generated by this strategy expire (LTM TTL).
	//
	// Minimum 31 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-recordexpirydurationdays
	//
	RecordExpiryDurationDays *float64 `field:"optional" json:"recordExpiryDurationDays" yaml:"recordExpiryDurationDays"`
	// Specifies the record expiry (LTM TTL) for a strategy.
	//
	// Provide RecordExpiryDurationDays to configure the duration after which records expire. Omit the specification entirely to leave records without an expiry on create, or to clear a previously configured expiry on update (the handler derives the removal from the previous and desired state).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-recordexpiryspecification
	//
	RecordExpirySpecification interface{} `field:"optional" json:"recordExpirySpecification" yaml:"recordExpirySpecification"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-reflectionconfiguration
	//
	ReflectionConfiguration interface{} `field:"optional" json:"reflectionConfiguration" yaml:"reflectionConfiguration"`
	// Status of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Unique identifier for the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// Type of memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Last update timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-episodicmemorystrategy.html#cfn-bedrockagentcore-memory-episodicmemorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}


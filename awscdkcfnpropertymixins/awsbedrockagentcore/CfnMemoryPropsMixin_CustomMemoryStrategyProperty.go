package awsbedrockagentcore


// The memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   customMemoryStrategyProperty := &CustomMemoryStrategyProperty{
//   	Configuration: &CustomConfigurationInputProperty{
//   		EpisodicOverride: &EpisodicOverrideProperty{
//   			Consolidation: &EpisodicOverrideConsolidationConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   			Extraction: &EpisodicOverrideExtractionConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   			Reflection: &EpisodicOverrideReflectionConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				MemoryRecordSchema: &MemoryRecordSchemaProperty{
//   					MetadataSchema: []interface{}{
//   						&MetadataSchemaEntryProperty{
//   							ExtractionConfig: &ExtractionConfigProperty{
//   								LlmExtractionConfig: &LlmExtractionConfigProperty{
//   									Definition: jsii.String("definition"),
//   									LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   									Validation: &ValidationProperty{
//   										NumberValidation: &NumberValidationProperty{
//   											MaxValue: jsii.Number(123),
//   											MinValue: jsii.Number(123),
//   										},
//   										StringListValidation: &StringListValidationProperty{
//   											AllowedValues: []*string{
//   												jsii.String("allowedValues"),
//   											},
//   											MaxItems: jsii.Number(123),
//   										},
//   										StringValidation: &StringValidationProperty{
//   											AllowedValues: []*string{
//   												jsii.String("allowedValues"),
//   											},
//   										},
//   									},
//   								},
//   							},
//   							ExtractionType: jsii.String("extractionType"),
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				ModelId: jsii.String("modelId"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   			},
//   		},
//   		SelfManagedConfiguration: &SelfManagedConfigurationProperty{
//   			HistoricalContextWindowSize: jsii.Number(123),
//   			InvocationConfiguration: &InvocationConfigurationInputProperty{
//   				PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   				TopicArn: jsii.String("topicArn"),
//   			},
//   			TriggerConditions: []interface{}{
//   				&TriggerConditionInputProperty{
//   					MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   						MessageCount: jsii.Number(123),
//   					},
//   					TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   						IdleSessionTimeout: jsii.Number(123),
//   					},
//   					TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   						TokenCount: jsii.Number(123),
//   					},
//   				},
//   			},
//   		},
//   		SemanticOverride: &SemanticOverrideProperty{
//   			Consolidation: &SemanticOverrideConsolidationConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   			Extraction: &SemanticOverrideExtractionConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   		},
//   		SummaryOverride: &SummaryOverrideProperty{
//   			Consolidation: &SummaryOverrideConsolidationConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   		},
//   		UserPreferenceOverride: &UserPreferenceOverrideProperty{
//   			Consolidation: &UserPreferenceOverrideConsolidationConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   			Extraction: &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   				AppendToPrompt: jsii.String("appendToPrompt"),
//   				ModelId: jsii.String("modelId"),
//   			},
//   		},
//   	},
//   	CreatedAt: jsii.String("createdAt"),
//   	Description: jsii.String("description"),
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
//   				ExtractionType: jsii.String("extractionType"),
//   				Key: jsii.String("key"),
//   				Type: jsii.String("type"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
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
//   	Status: jsii.String("status"),
//   	StrategyId: jsii.String("strategyId"),
//   	Type: jsii.String("type"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html
//
type CfnMemoryPropsMixin_CustomMemoryStrategyProperty struct {
	// The memory strategy configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// The memory strategy description.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-memoryrecordschema
	//
	MemoryRecordSchema interface{} `field:"optional" json:"memoryRecordSchema" yaml:"memoryRecordSchema"`
	// The memory strategy name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The memory strategy namespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-namespacetemplates
	//
	NamespaceTemplates *[]*string `field:"optional" json:"namespaceTemplates" yaml:"namespaceTemplates"`
	// The duration, in days, after which long-term memory records generated by this strategy expire (LTM TTL).
	//
	// Minimum 31 days.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-recordexpirydurationdays
	//
	RecordExpiryDurationDays *float64 `field:"optional" json:"recordExpiryDurationDays" yaml:"recordExpiryDurationDays"`
	// Specifies the record expiry (LTM TTL) for a strategy.
	//
	// Provide RecordExpiryDurationDays to configure the duration after which records expire. Omit the specification entirely to leave records without an expiry on create, or to clear a previously configured expiry on update (the handler derives the removal from the previous and desired state).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-recordexpiryspecification
	//
	RecordExpirySpecification interface{} `field:"optional" json:"recordExpirySpecification" yaml:"recordExpirySpecification"`
	// The memory strategy status.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// The memory strategy ID.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// The memory strategy type.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The memory strategy update date and time.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}


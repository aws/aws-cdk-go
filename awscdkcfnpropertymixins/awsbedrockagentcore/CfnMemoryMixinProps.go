package awsbedrockagentcore


// Properties for CfnMemoryPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   cfnMemoryMixinProps := &CfnMemoryMixinProps{
//   	Description: jsii.String("description"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	EventExpiryDuration: jsii.Number(123),
//   	IndexedKeys: []interface{}{
//   		&IndexedKeyProperty{
//   			Key: jsii.String("key"),
//   			Type: jsii.String("type"),
//   		},
//   	},
//   	MemoryExecutionRoleArn: jsii.String("memoryExecutionRoleArn"),
//   	MemoryStrategies: []interface{}{
//   		&MemoryStrategyProperty{
//   			CustomMemoryStrategy: &CustomMemoryStrategyProperty{
//   				Configuration: &CustomConfigurationInputProperty{
//   					EpisodicOverride: &EpisodicOverrideProperty{
//   						Consolidation: &EpisodicOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Extraction: &EpisodicOverrideExtractionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Reflection: &EpisodicOverrideReflectionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							MemoryRecordSchema: &MemoryRecordSchemaProperty{
//   								MetadataSchema: []interface{}{
//   									&MetadataSchemaEntryProperty{
//   										ExtractionConfig: &ExtractionConfigProperty{
//   											LlmExtractionConfig: &LlmExtractionConfigProperty{
//   												Definition: jsii.String("definition"),
//   												LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   												Validation: &ValidationProperty{
//   													NumberValidation: &NumberValidationProperty{
//   														MaxValue: jsii.Number(123),
//   														MinValue: jsii.Number(123),
//   													},
//   													StringListValidation: &StringListValidationProperty{
//   														AllowedValues: []*string{
//   															jsii.String("allowedValues"),
//   														},
//   														MaxItems: jsii.Number(123),
//   													},
//   													StringValidation: &StringValidationProperty{
//   														AllowedValues: []*string{
//   															jsii.String("allowedValues"),
//   														},
//   													},
//   												},
//   											},
//   										},
//   										Key: jsii.String("key"),
//   										Type: jsii.String("type"),
//   									},
//   								},
//   							},
//   							ModelId: jsii.String("modelId"),
//   							Namespaces: []*string{
//   								jsii.String("namespaces"),
//   							},
//   							NamespaceTemplates: []*string{
//   								jsii.String("namespaceTemplates"),
//   							},
//   						},
//   					},
//   					SelfManagedConfiguration: &SelfManagedConfigurationProperty{
//   						HistoricalContextWindowSize: jsii.Number(123),
//   						InvocationConfiguration: &InvocationConfigurationInputProperty{
//   							PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   							TopicArn: jsii.String("topicArn"),
//   						},
//   						TriggerConditions: []interface{}{
//   							&TriggerConditionInputProperty{
//   								MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   									MessageCount: jsii.Number(123),
//   								},
//   								TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   									IdleSessionTimeout: jsii.Number(123),
//   								},
//   								TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   									TokenCount: jsii.Number(123),
//   								},
//   							},
//   						},
//   					},
//   					SemanticOverride: &SemanticOverrideProperty{
//   						Consolidation: &SemanticOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Extraction: &SemanticOverrideExtractionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   					},
//   					SummaryOverride: &SummaryOverrideProperty{
//   						Consolidation: &SummaryOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   					},
//   					UserPreferenceOverride: &UserPreferenceOverrideProperty{
//   						Consolidation: &UserPreferenceOverrideConsolidationConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   						Extraction: &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   							AppendToPrompt: jsii.String("appendToPrompt"),
//   							ModelId: jsii.String("modelId"),
//   						},
//   					},
//   				},
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			EpisodicMemoryStrategy: &EpisodicMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				ReflectionConfiguration: &EpisodicReflectionConfigurationInputProperty{
//   					MemoryRecordSchema: &MemoryRecordSchemaProperty{
//   						MetadataSchema: []interface{}{
//   							&MetadataSchemaEntryProperty{
//   								ExtractionConfig: &ExtractionConfigProperty{
//   									LlmExtractionConfig: &LlmExtractionConfigProperty{
//   										Definition: jsii.String("definition"),
//   										LlmExtractionInstruction: jsii.String("llmExtractionInstruction"),
//   										Validation: &ValidationProperty{
//   											NumberValidation: &NumberValidationProperty{
//   												MaxValue: jsii.Number(123),
//   												MinValue: jsii.Number(123),
//   											},
//   											StringListValidation: &StringListValidationProperty{
//   												AllowedValues: []*string{
//   													jsii.String("allowedValues"),
//   												},
//   												MaxItems: jsii.Number(123),
//   											},
//   											StringValidation: &StringValidationProperty{
//   												AllowedValues: []*string{
//   													jsii.String("allowedValues"),
//   												},
//   											},
//   										},
//   									},
//   								},
//   								Key: jsii.String("key"),
//   								Type: jsii.String("type"),
//   							},
//   						},
//   					},
//   					Namespaces: []*string{
//   						jsii.String("namespaces"),
//   					},
//   					NamespaceTemplates: []*string{
//   						jsii.String("namespaceTemplates"),
//   					},
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			SemanticMemoryStrategy: &SemanticMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			SummaryMemoryStrategy: &SummaryMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   			UserPreferenceMemoryStrategy: &UserPreferenceMemoryStrategyProperty{
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   							Key: jsii.String("key"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   				},
//   				Name: jsii.String("name"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
//   			},
//   		},
//   	},
//   	Name: jsii.String("name"),
//   	StreamDeliveryResources: &StreamDeliveryResourcesProperty{
//   		Resources: []interface{}{
//   			&StreamDeliveryResourceProperty{
//   				Kinesis: &KinesisResourceProperty{
//   					ContentConfigurations: []interface{}{
//   						&ContentConfigurationProperty{
//   							Level: jsii.String("level"),
//   							Type: jsii.String("type"),
//   						},
//   					},
//   					DataStreamArn: jsii.String("dataStreamArn"),
//   				},
//   			},
//   		},
//   	},
//   	Tags: map[string]*string{
//   		"tagsKey": jsii.String("tags"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html
//
type CfnMemoryMixinProps struct {
	// Description of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The memory encryption key Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The event expiry configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-eventexpiryduration
	//
	EventExpiryDuration *float64 `field:"optional" json:"eventExpiryDuration" yaml:"eventExpiryDuration"`
	// List of indexed keys for the memory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-indexedkeys
	//
	IndexedKeys interface{} `field:"optional" json:"indexedKeys" yaml:"indexedKeys"`
	// The memory role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-memoryexecutionrolearn
	//
	MemoryExecutionRoleArn *string `field:"optional" json:"memoryExecutionRoleArn" yaml:"memoryExecutionRoleArn"`
	// The memory strategies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-memorystrategies
	//
	MemoryStrategies interface{} `field:"optional" json:"memoryStrategies" yaml:"memoryStrategies"`
	// The memory name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-streamdeliveryresources
	//
	StreamDeliveryResources interface{} `field:"optional" json:"streamDeliveryResources" yaml:"streamDeliveryResources"`
	// The tags for the resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


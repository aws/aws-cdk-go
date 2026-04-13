package awsbedrockagentcore


// Properties for defining a `CfnMemory`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   cfnMemoryProps := &CfnMemoryProps{
//   	EventExpiryDuration: jsii.Number(123),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
//   	Description: jsii.String("description"),
//   	EncryptionKeyArn: jsii.String("encryptionKeyArn"),
//   	MemoryExecutionRoleArn: jsii.String("memoryExecutionRoleArn"),
//   	MemoryStrategies: []interface{}{
//   		&MemoryStrategyProperty{
//   			CustomMemoryStrategy: &CustomMemoryStrategyProperty{
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
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
//   							ModelId: jsii.String("modelId"),
//
//   							// the properties below are optional
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
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
//   				},
//   				NamespaceTemplates: []*string{
//   					jsii.String("namespaceTemplates"),
//   				},
//   				ReflectionConfiguration: &EpisodicReflectionConfigurationInputProperty{
//   					Namespaces: []*string{
//   						jsii.String("namespaces"),
//   					},
//
//   					// the properties below are optional
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
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   				Name: jsii.String("name"),
//
//   				// the properties below are optional
//   				CreatedAt: jsii.String("createdAt"),
//   				Description: jsii.String("description"),
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
//   	StreamDeliveryResources: &StreamDeliveryResourcesProperty{
//   		Resources: []interface{}{
//   			&StreamDeliveryResourceProperty{
//   				Kinesis: &KinesisResourceProperty{
//   					ContentConfigurations: []interface{}{
//   						&ContentConfigurationProperty{
//   							Type: jsii.String("type"),
//
//   							// the properties below are optional
//   							Level: jsii.String("level"),
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
type CfnMemoryProps struct {
	// The event expiry configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-eventexpiryduration
	//
	EventExpiryDuration *float64 `field:"required" json:"eventExpiryDuration" yaml:"eventExpiryDuration"`
	// The memory name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The memory encryption key Amazon Resource Name (ARN).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// The memory role ARN.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-memoryexecutionrolearn
	//
	MemoryExecutionRoleArn *string `field:"optional" json:"memoryExecutionRoleArn" yaml:"memoryExecutionRoleArn"`
	// The memory strategies.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-memorystrategies
	//
	MemoryStrategies interface{} `field:"optional" json:"memoryStrategies" yaml:"memoryStrategies"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-streamdeliveryresources
	//
	StreamDeliveryResources interface{} `field:"optional" json:"streamDeliveryResources" yaml:"streamDeliveryResources"`
	// The tags for the resources.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


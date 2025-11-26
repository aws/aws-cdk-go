package previewawsbedrockagentcoremixins


// The memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   memoryStrategyProperty := &MemoryStrategyProperty{
//   	CustomMemoryStrategy: &CustomMemoryStrategyProperty{
//   		Configuration: &CustomConfigurationInputProperty{
//   			SelfManagedConfiguration: &SelfManagedConfigurationProperty{
//   				HistoricalContextWindowSize: jsii.Number(123),
//   				InvocationConfiguration: &InvocationConfigurationInputProperty{
//   					PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   					TopicArn: jsii.String("topicArn"),
//   				},
//   				TriggerConditions: []interface{}{
//   					&TriggerConditionInputProperty{
//   						MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   							MessageCount: jsii.Number(123),
//   						},
//   						TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   							IdleSessionTimeout: jsii.Number(123),
//   						},
//   						TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   							TokenCount: jsii.Number(123),
//   						},
//   					},
//   				},
//   			},
//   			SemanticOverride: &SemanticOverrideProperty{
//   				Consolidation: &SemanticOverrideConsolidationConfigurationInputProperty{
//   					AppendToPrompt: jsii.String("appendToPrompt"),
//   					ModelId: jsii.String("modelId"),
//   				},
//   				Extraction: &SemanticOverrideExtractionConfigurationInputProperty{
//   					AppendToPrompt: jsii.String("appendToPrompt"),
//   					ModelId: jsii.String("modelId"),
//   				},
//   			},
//   			SummaryOverride: &SummaryOverrideProperty{
//   				Consolidation: &SummaryOverrideConsolidationConfigurationInputProperty{
//   					AppendToPrompt: jsii.String("appendToPrompt"),
//   					ModelId: jsii.String("modelId"),
//   				},
//   			},
//   			UserPreferenceOverride: &UserPreferenceOverrideProperty{
//   				Consolidation: &UserPreferenceOverrideConsolidationConfigurationInputProperty{
//   					AppendToPrompt: jsii.String("appendToPrompt"),
//   					ModelId: jsii.String("modelId"),
//   				},
//   				Extraction: &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   					AppendToPrompt: jsii.String("appendToPrompt"),
//   					ModelId: jsii.String("modelId"),
//   				},
//   			},
//   		},
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	SemanticMemoryStrategy: &SemanticMemoryStrategyProperty{
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	SummaryMemoryStrategy: &SummaryMemoryStrategyProperty{
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	UserPreferenceMemoryStrategy: &UserPreferenceMemoryStrategyProperty{
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
//   		Name: jsii.String("name"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html
//
type CfnMemoryPropsMixin_MemoryStrategyProperty struct {
	// The memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-custommemorystrategy
	//
	CustomMemoryStrategy interface{} `field:"optional" json:"customMemoryStrategy" yaml:"customMemoryStrategy"`
	// The memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-semanticmemorystrategy
	//
	SemanticMemoryStrategy interface{} `field:"optional" json:"semanticMemoryStrategy" yaml:"semanticMemoryStrategy"`
	// The memory strategy summary.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-summarymemorystrategy
	//
	SummaryMemoryStrategy interface{} `field:"optional" json:"summaryMemoryStrategy" yaml:"summaryMemoryStrategy"`
	// The memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-userpreferencememorystrategy
	//
	UserPreferenceMemoryStrategy interface{} `field:"optional" json:"userPreferenceMemoryStrategy" yaml:"userPreferenceMemoryStrategy"`
}


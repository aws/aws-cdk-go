package awsbedrockagentcore


// The memory strategy.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customMemoryStrategyProperty := &CustomMemoryStrategyProperty{
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
//   				ModelId: jsii.String("modelId"),
//
//   				// the properties below are optional
//   				Namespaces: []*string{
//   					jsii.String("namespaces"),
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
//   	Namespaces: []*string{
//   		jsii.String("namespaces"),
//   	},
//   	Status: jsii.String("status"),
//   	StrategyId: jsii.String("strategyId"),
//   	Type: jsii.String("type"),
//   	UpdatedAt: jsii.String("updatedAt"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html
//
type CfnMemory_CustomMemoryStrategyProperty struct {
	// The memory strategy name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
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
	// The memory strategy namespaces.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
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


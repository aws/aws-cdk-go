package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   memoryStrategyProperty := &MemoryStrategyProperty{
//   	CustomMemoryStrategy: &CustomMemoryStrategyProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		Configuration: &CustomConfigurationInputProperty{
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
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	SemanticMemoryStrategy: &SemanticMemoryStrategyProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	SummaryMemoryStrategy: &SummaryMemoryStrategyProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
//   		Namespaces: []*string{
//   			jsii.String("namespaces"),
//   		},
//   		Status: jsii.String("status"),
//   		StrategyId: jsii.String("strategyId"),
//   		Type: jsii.String("type"),
//   		UpdatedAt: jsii.String("updatedAt"),
//   	},
//   	UserPreferenceMemoryStrategy: &UserPreferenceMemoryStrategyProperty{
//   		Name: jsii.String("name"),
//
//   		// the properties below are optional
//   		CreatedAt: jsii.String("createdAt"),
//   		Description: jsii.String("description"),
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
type CfnMemory_MemoryStrategyProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-custommemorystrategy
	//
	CustomMemoryStrategy interface{} `field:"optional" json:"customMemoryStrategy" yaml:"customMemoryStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-semanticmemorystrategy
	//
	SemanticMemoryStrategy interface{} `field:"optional" json:"semanticMemoryStrategy" yaml:"semanticMemoryStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-summarymemorystrategy
	//
	SummaryMemoryStrategy interface{} `field:"optional" json:"summaryMemoryStrategy" yaml:"summaryMemoryStrategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-memorystrategy.html#cfn-bedrockagentcore-memory-memorystrategy-userpreferencememorystrategy
	//
	UserPreferenceMemoryStrategy interface{} `field:"optional" json:"userPreferenceMemoryStrategy" yaml:"userPreferenceMemoryStrategy"`
}


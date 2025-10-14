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
//   				Status: jsii.String("status"),
//   				StrategyId: jsii.String("strategyId"),
//   				Type: jsii.String("type"),
//   				UpdatedAt: jsii.String("updatedAt"),
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
	// Duration in days until memory events expire.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-eventexpiryduration
	//
	EventExpiryDuration *float64 `field:"required" json:"eventExpiryDuration" yaml:"eventExpiryDuration"`
	// Name of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Description of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// ARN format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-encryptionkeyarn
	//
	EncryptionKeyArn *string `field:"optional" json:"encryptionKeyArn" yaml:"encryptionKeyArn"`
	// ARN format.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-memoryexecutionrolearn
	//
	MemoryExecutionRoleArn *string `field:"optional" json:"memoryExecutionRoleArn" yaml:"memoryExecutionRoleArn"`
	// List of memory strategies attached to this memory.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-memorystrategies
	//
	MemoryStrategies interface{} `field:"optional" json:"memoryStrategies" yaml:"memoryStrategies"`
	// A map of tag keys and values.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-bedrockagentcore-memory.html#cfn-bedrockagentcore-memory-tags
	//
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
}


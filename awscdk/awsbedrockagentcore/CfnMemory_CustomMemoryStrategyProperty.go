package awsbedrockagentcore


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
	// Name of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-configuration
	//
	Configuration interface{} `field:"optional" json:"configuration" yaml:"configuration"`
	// Creation timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-createdat
	//
	CreatedAt *string `field:"optional" json:"createdAt" yaml:"createdAt"`
	// Description of the Memory resource.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-description
	//
	Description *string `field:"optional" json:"description" yaml:"description"`
	// List of namespaces for memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-namespaces
	//
	Namespaces *[]*string `field:"optional" json:"namespaces" yaml:"namespaces"`
	// Status of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-status
	//
	Status *string `field:"optional" json:"status" yaml:"status"`
	// Unique identifier for the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-strategyid
	//
	StrategyId *string `field:"optional" json:"strategyId" yaml:"strategyId"`
	// Type of memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-type
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Last update timestamp of the memory strategy.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-custommemorystrategy.html#cfn-bedrockagentcore-memory-custommemorystrategy-updatedat
	//
	UpdatedAt *string `field:"optional" json:"updatedAt" yaml:"updatedAt"`
}


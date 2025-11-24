package mixinsawsbedrockagentcore


// The memory configuration input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   customConfigurationInputProperty := &CustomConfigurationInputProperty{
//   	SelfManagedConfiguration: &SelfManagedConfigurationProperty{
//   		HistoricalContextWindowSize: jsii.Number(123),
//   		InvocationConfiguration: &InvocationConfigurationInputProperty{
//   			PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   			TopicArn: jsii.String("topicArn"),
//   		},
//   		TriggerConditions: []interface{}{
//   			&TriggerConditionInputProperty{
//   				MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   					MessageCount: jsii.Number(123),
//   				},
//   				TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   					IdleSessionTimeout: jsii.Number(123),
//   				},
//   				TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   					TokenCount: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	SemanticOverride: &SemanticOverrideProperty{
//   		Consolidation: &SemanticOverrideConsolidationConfigurationInputProperty{
//   			AppendToPrompt: jsii.String("appendToPrompt"),
//   			ModelId: jsii.String("modelId"),
//   		},
//   		Extraction: &SemanticOverrideExtractionConfigurationInputProperty{
//   			AppendToPrompt: jsii.String("appendToPrompt"),
//   			ModelId: jsii.String("modelId"),
//   		},
//   	},
//   	SummaryOverride: &SummaryOverrideProperty{
//   		Consolidation: &SummaryOverrideConsolidationConfigurationInputProperty{
//   			AppendToPrompt: jsii.String("appendToPrompt"),
//   			ModelId: jsii.String("modelId"),
//   		},
//   	},
//   	UserPreferenceOverride: &UserPreferenceOverrideProperty{
//   		Consolidation: &UserPreferenceOverrideConsolidationConfigurationInputProperty{
//   			AppendToPrompt: jsii.String("appendToPrompt"),
//   			ModelId: jsii.String("modelId"),
//   		},
//   		Extraction: &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   			AppendToPrompt: jsii.String("appendToPrompt"),
//   			ModelId: jsii.String("modelId"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html
//
type CfnMemoryPropsMixin_CustomConfigurationInputProperty struct {
	// The custom configuration input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-selfmanagedconfiguration
	//
	SelfManagedConfiguration interface{} `field:"optional" json:"selfManagedConfiguration" yaml:"selfManagedConfiguration"`
	// The memory override configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-semanticoverride
	//
	SemanticOverride interface{} `field:"optional" json:"semanticOverride" yaml:"semanticOverride"`
	// The memory configuration override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-summaryoverride
	//
	SummaryOverride interface{} `field:"optional" json:"summaryOverride" yaml:"summaryOverride"`
	// The memory user preference override.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-userpreferenceoverride
	//
	UserPreferenceOverride interface{} `field:"optional" json:"userPreferenceOverride" yaml:"userPreferenceOverride"`
}


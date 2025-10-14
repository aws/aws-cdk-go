package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   customConfigurationInputProperty := &CustomConfigurationInputProperty{
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
type CfnMemory_CustomConfigurationInputProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-semanticoverride
	//
	SemanticOverride interface{} `field:"optional" json:"semanticOverride" yaml:"semanticOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-summaryoverride
	//
	SummaryOverride interface{} `field:"optional" json:"summaryOverride" yaml:"summaryOverride"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-customconfigurationinput.html#cfn-bedrockagentcore-memory-customconfigurationinput-userpreferenceoverride
	//
	UserPreferenceOverride interface{} `field:"optional" json:"userPreferenceOverride" yaml:"userPreferenceOverride"`
}


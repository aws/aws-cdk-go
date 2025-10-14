package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   userPreferenceOverrideProperty := &UserPreferenceOverrideProperty{
//   	Consolidation: &UserPreferenceOverrideConsolidationConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   	Extraction: &UserPreferenceOverrideExtractionConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferenceoverride.html
//
type CfnMemory_UserPreferenceOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferenceoverride.html#cfn-bedrockagentcore-memory-userpreferenceoverride-consolidation
	//
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-userpreferenceoverride.html#cfn-bedrockagentcore-memory-userpreferenceoverride-extraction
	//
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
}


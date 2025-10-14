package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   summaryOverrideProperty := &SummaryOverrideProperty{
//   	Consolidation: &SummaryOverrideConsolidationConfigurationInputProperty{
//   		AppendToPrompt: jsii.String("appendToPrompt"),
//   		ModelId: jsii.String("modelId"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summaryoverride.html
//
type CfnMemory_SummaryOverrideProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-summaryoverride.html#cfn-bedrockagentcore-memory-summaryoverride-consolidation
	//
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
}


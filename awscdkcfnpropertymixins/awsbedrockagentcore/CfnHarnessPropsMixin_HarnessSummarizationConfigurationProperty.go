package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessSummarizationConfigurationProperty := &HarnessSummarizationConfigurationProperty{
//   	PreserveRecentMessages: jsii.Number(123),
//   	SummarizationSystemPrompt: jsii.String("summarizationSystemPrompt"),
//   	SummaryRatio: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesssummarizationconfiguration.html
//
type CfnHarnessPropsMixin_HarnessSummarizationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesssummarizationconfiguration.html#cfn-bedrockagentcore-harness-harnesssummarizationconfiguration-preserverecentmessages
	//
	PreserveRecentMessages *float64 `field:"optional" json:"preserveRecentMessages" yaml:"preserveRecentMessages"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesssummarizationconfiguration.html#cfn-bedrockagentcore-harness-harnesssummarizationconfiguration-summarizationsystemprompt
	//
	SummarizationSystemPrompt *string `field:"optional" json:"summarizationSystemPrompt" yaml:"summarizationSystemPrompt"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesssummarizationconfiguration.html#cfn-bedrockagentcore-harness-harnesssummarizationconfiguration-summaryratio
	//
	SummaryRatio *float64 `field:"optional" json:"summaryRatio" yaml:"summaryRatio"`
}


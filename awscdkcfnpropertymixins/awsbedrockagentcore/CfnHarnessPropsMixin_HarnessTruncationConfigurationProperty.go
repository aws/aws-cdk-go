package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessTruncationConfigurationProperty := &HarnessTruncationConfigurationProperty{
//   	Config: &HarnessTruncationStrategyConfigurationProperty{
//   		SlidingWindow: &HarnessSlidingWindowConfigurationProperty{
//   			MessagesCount: jsii.Number(123),
//   		},
//   		Summarization: &HarnessSummarizationConfigurationProperty{
//   			PreserveRecentMessages: jsii.Number(123),
//   			SummarizationSystemPrompt: jsii.String("summarizationSystemPrompt"),
//   			SummaryRatio: jsii.Number(123),
//   		},
//   	},
//   	Strategy: jsii.String("strategy"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationconfiguration.html
//
type CfnHarnessPropsMixin_HarnessTruncationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationconfiguration.html#cfn-bedrockagentcore-harness-harnesstruncationconfiguration-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationconfiguration.html#cfn-bedrockagentcore-harness-harnesstruncationconfiguration-strategy
	//
	Strategy *string `field:"optional" json:"strategy" yaml:"strategy"`
}


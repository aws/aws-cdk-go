package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessTruncationConfigurationProperty := &HarnessTruncationConfigurationProperty{
//   	Strategy: jsii.String("strategy"),
//
//   	// the properties below are optional
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
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationconfiguration.html
//
type CfnHarness_HarnessTruncationConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationconfiguration.html#cfn-bedrockagentcore-harness-harnesstruncationconfiguration-strategy
	//
	Strategy *string `field:"required" json:"strategy" yaml:"strategy"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationconfiguration.html#cfn-bedrockagentcore-harness-harnesstruncationconfiguration-config
	//
	Config interface{} `field:"optional" json:"config" yaml:"config"`
}


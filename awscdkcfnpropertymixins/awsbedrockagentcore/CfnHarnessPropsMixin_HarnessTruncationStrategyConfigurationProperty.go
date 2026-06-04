package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkcfnpropertymixins"
//
//   harnessTruncationStrategyConfigurationProperty := &HarnessTruncationStrategyConfigurationProperty{
//   	SlidingWindow: &HarnessSlidingWindowConfigurationProperty{
//   		MessagesCount: jsii.Number(123),
//   	},
//   	Summarization: &HarnessSummarizationConfigurationProperty{
//   		PreserveRecentMessages: jsii.Number(123),
//   		SummarizationSystemPrompt: jsii.String("summarizationSystemPrompt"),
//   		SummaryRatio: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationstrategyconfiguration.html
//
type CfnHarnessPropsMixin_HarnessTruncationStrategyConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationstrategyconfiguration.html#cfn-bedrockagentcore-harness-harnesstruncationstrategyconfiguration-slidingwindow
	//
	SlidingWindow interface{} `field:"optional" json:"slidingWindow" yaml:"slidingWindow"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnesstruncationstrategyconfiguration.html#cfn-bedrockagentcore-harness-harnesstruncationstrategyconfiguration-summarization
	//
	Summarization interface{} `field:"optional" json:"summarization" yaml:"summarization"`
}


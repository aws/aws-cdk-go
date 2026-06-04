package awsbedrockagentcore


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   harnessSlidingWindowConfigurationProperty := &HarnessSlidingWindowConfigurationProperty{
//   	MessagesCount: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessslidingwindowconfiguration.html
//
type CfnHarness_HarnessSlidingWindowConfigurationProperty struct {
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-harness-harnessslidingwindowconfiguration.html#cfn-bedrockagentcore-harness-harnessslidingwindowconfiguration-messagescount
	//
	MessagesCount *float64 `field:"optional" json:"messagesCount" yaml:"messagesCount"`
}


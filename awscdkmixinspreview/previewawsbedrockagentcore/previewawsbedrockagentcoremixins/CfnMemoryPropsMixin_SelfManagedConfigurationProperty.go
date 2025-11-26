package previewawsbedrockagentcoremixins


// The self managed configuration.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   selfManagedConfigurationProperty := &SelfManagedConfigurationProperty{
//   	HistoricalContextWindowSize: jsii.Number(123),
//   	InvocationConfiguration: &InvocationConfigurationInputProperty{
//   		PayloadDeliveryBucketName: jsii.String("payloadDeliveryBucketName"),
//   		TopicArn: jsii.String("topicArn"),
//   	},
//   	TriggerConditions: []interface{}{
//   		&TriggerConditionInputProperty{
//   			MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   				MessageCount: jsii.Number(123),
//   			},
//   			TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   				IdleSessionTimeout: jsii.Number(123),
//   			},
//   			TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   				TokenCount: jsii.Number(123),
//   			},
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-selfmanagedconfiguration.html
//
type CfnMemoryPropsMixin_SelfManagedConfigurationProperty struct {
	// The memory configuration for self managed.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-selfmanagedconfiguration.html#cfn-bedrockagentcore-memory-selfmanagedconfiguration-historicalcontextwindowsize
	//
	HistoricalContextWindowSize *float64 `field:"optional" json:"historicalContextWindowSize" yaml:"historicalContextWindowSize"`
	// The self managed configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-selfmanagedconfiguration.html#cfn-bedrockagentcore-memory-selfmanagedconfiguration-invocationconfiguration
	//
	InvocationConfiguration interface{} `field:"optional" json:"invocationConfiguration" yaml:"invocationConfiguration"`
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-selfmanagedconfiguration.html#cfn-bedrockagentcore-memory-selfmanagedconfiguration-triggerconditions
	//
	TriggerConditions interface{} `field:"optional" json:"triggerConditions" yaml:"triggerConditions"`
}


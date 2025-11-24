package mixinsawsbedrockagentcore


// The memory trigger condition input.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   triggerConditionInputProperty := &TriggerConditionInputProperty{
//   	MessageBasedTrigger: &MessageBasedTriggerInputProperty{
//   		MessageCount: jsii.Number(123),
//   	},
//   	TimeBasedTrigger: &TimeBasedTriggerInputProperty{
//   		IdleSessionTimeout: jsii.Number(123),
//   	},
//   	TokenBasedTrigger: &TokenBasedTriggerInputProperty{
//   		TokenCount: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-triggerconditioninput.html
//
type CfnMemoryPropsMixin_TriggerConditionInputProperty struct {
	// The memory trigger condition input for the message based trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-triggerconditioninput.html#cfn-bedrockagentcore-memory-triggerconditioninput-messagebasedtrigger
	//
	MessageBasedTrigger interface{} `field:"optional" json:"messageBasedTrigger" yaml:"messageBasedTrigger"`
	// The memory trigger condition input.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-triggerconditioninput.html#cfn-bedrockagentcore-memory-triggerconditioninput-timebasedtrigger
	//
	TimeBasedTrigger interface{} `field:"optional" json:"timeBasedTrigger" yaml:"timeBasedTrigger"`
	// The trigger condition information for a token based trigger.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrockagentcore-memory-triggerconditioninput.html#cfn-bedrockagentcore-memory-triggerconditioninput-tokenbasedtrigger
	//
	TokenBasedTrigger interface{} `field:"optional" json:"tokenBasedTrigger" yaml:"tokenBasedTrigger"`
}


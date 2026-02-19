package awsconnect


// After Contact Work configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   afterContactWorkConfigPerChannelProperty := &AfterContactWorkConfigPerChannelProperty{
//   	AfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   	},
//   	Channel: jsii.String("channel"),
//
//   	// the properties below are optional
//   	AgentFirstCallbackAfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html
//
type CfnUser_AfterContactWorkConfigPerChannelProperty struct {
	// After Contact Work configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html#cfn-connect-user-aftercontactworkconfigperchannel-aftercontactworkconfig
	//
	AfterContactWorkConfig interface{} `field:"required" json:"afterContactWorkConfig" yaml:"afterContactWorkConfig"`
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html#cfn-connect-user-aftercontactworkconfigperchannel-channel
	//
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// After Contact Work configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html#cfn-connect-user-aftercontactworkconfigperchannel-agentfirstcallbackaftercontactworkconfig
	//
	AgentFirstCallbackAfterContactWorkConfig interface{} `field:"optional" json:"agentFirstCallbackAfterContactWorkConfig" yaml:"agentFirstCallbackAfterContactWorkConfig"`
}


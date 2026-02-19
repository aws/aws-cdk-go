package previewawsconnectmixins


// After Contact Work configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   afterContactWorkConfigPerChannelProperty := &AfterContactWorkConfigPerChannelProperty{
//   	AfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   	},
//   	AgentFirstCallbackAfterContactWorkConfig: &AfterContactWorkConfigProperty{
//   		AfterContactWorkTimeLimit: jsii.Number(123),
//   	},
//   	Channel: jsii.String("channel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html
//
type CfnUserPropsMixin_AfterContactWorkConfigPerChannelProperty struct {
	// After Contact Work configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html#cfn-connect-user-aftercontactworkconfigperchannel-aftercontactworkconfig
	//
	AfterContactWorkConfig interface{} `field:"optional" json:"afterContactWorkConfig" yaml:"afterContactWorkConfig"`
	// After Contact Work configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html#cfn-connect-user-aftercontactworkconfigperchannel-agentfirstcallbackaftercontactworkconfig
	//
	AgentFirstCallbackAfterContactWorkConfig interface{} `field:"optional" json:"agentFirstCallbackAfterContactWorkConfig" yaml:"agentFirstCallbackAfterContactWorkConfig"`
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-aftercontactworkconfigperchannel.html#cfn-connect-user-aftercontactworkconfigperchannel-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
}


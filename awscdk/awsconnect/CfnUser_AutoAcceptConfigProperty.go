package awsconnect


// Auto-accept configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   autoAcceptConfigProperty := &AutoAcceptConfigProperty{
//   	AutoAccept: jsii.Boolean(false),
//   	Channel: jsii.String("channel"),
//
//   	// the properties below are optional
//   	AgentFirstCallbackAutoAccept: jsii.Boolean(false),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html
//
type CfnUser_AutoAcceptConfigProperty struct {
	// The Auto accept setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html#cfn-connect-user-autoacceptconfig-autoaccept
	//
	AutoAccept interface{} `field:"required" json:"autoAccept" yaml:"autoAccept"`
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html#cfn-connect-user-autoacceptconfig-channel
	//
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// The agent first callback auto accept setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html#cfn-connect-user-autoacceptconfig-agentfirstcallbackautoaccept
	//
	AgentFirstCallbackAutoAccept interface{} `field:"optional" json:"agentFirstCallbackAutoAccept" yaml:"agentFirstCallbackAutoAccept"`
}


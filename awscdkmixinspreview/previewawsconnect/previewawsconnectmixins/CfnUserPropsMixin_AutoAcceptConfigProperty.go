package previewawsconnectmixins


// Auto-accept configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   autoAcceptConfigProperty := &AutoAcceptConfigProperty{
//   	AgentFirstCallbackAutoAccept: jsii.Boolean(false),
//   	AutoAccept: jsii.Boolean(false),
//   	Channel: jsii.String("channel"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html
//
type CfnUserPropsMixin_AutoAcceptConfigProperty struct {
	// The agent first callback auto accept setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html#cfn-connect-user-autoacceptconfig-agentfirstcallbackautoaccept
	//
	AgentFirstCallbackAutoAccept interface{} `field:"optional" json:"agentFirstCallbackAutoAccept" yaml:"agentFirstCallbackAutoAccept"`
	// The Auto accept setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html#cfn-connect-user-autoacceptconfig-autoaccept
	//
	AutoAccept interface{} `field:"optional" json:"autoAccept" yaml:"autoAccept"`
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-autoacceptconfig.html#cfn-connect-user-autoacceptconfig-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
}


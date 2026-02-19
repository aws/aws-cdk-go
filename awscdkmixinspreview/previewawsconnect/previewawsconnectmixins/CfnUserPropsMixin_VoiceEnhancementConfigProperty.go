package previewawsconnectmixins


// Voice Enhancement configuration per channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   voiceEnhancementConfigProperty := &VoiceEnhancementConfigProperty{
//   	Channel: jsii.String("channel"),
//   	VoiceEnhancementMode: jsii.String("voiceEnhancementMode"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-voiceenhancementconfig.html
//
type CfnUserPropsMixin_VoiceEnhancementConfigProperty struct {
	// The channels that agents can handle in the Contact Control Panel (CCP).
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-voiceenhancementconfig.html#cfn-connect-user-voiceenhancementconfig-channel
	//
	Channel *string `field:"optional" json:"channel" yaml:"channel"`
	// The Voice Enhancement Mode setting.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connect-user-voiceenhancementconfig.html#cfn-connect-user-voiceenhancementconfig-voiceenhancementmode
	//
	VoiceEnhancementMode *string `field:"optional" json:"voiceEnhancementMode" yaml:"voiceEnhancementMode"`
}


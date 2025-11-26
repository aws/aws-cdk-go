package previewawspinpointmixins


// Specifies the message template to use for the message, for each type of channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   templateConfigurationProperty := &TemplateConfigurationProperty{
//   	EmailTemplate: &TemplateProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	PushTemplate: &TemplateProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	SmsTemplate: &TemplateProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   	VoiceTemplate: &TemplateProperty{
//   		Name: jsii.String("name"),
//   		Version: jsii.String("version"),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-templateconfiguration.html
//
type CfnCampaignPropsMixin_TemplateConfigurationProperty struct {
	// The email template to use for the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-templateconfiguration.html#cfn-pinpoint-campaign-templateconfiguration-emailtemplate
	//
	EmailTemplate interface{} `field:"optional" json:"emailTemplate" yaml:"emailTemplate"`
	// The push notification template to use for the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-templateconfiguration.html#cfn-pinpoint-campaign-templateconfiguration-pushtemplate
	//
	PushTemplate interface{} `field:"optional" json:"pushTemplate" yaml:"pushTemplate"`
	// The SMS template to use for the message.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-templateconfiguration.html#cfn-pinpoint-campaign-templateconfiguration-smstemplate
	//
	SmsTemplate interface{} `field:"optional" json:"smsTemplate" yaml:"smsTemplate"`
	// The voice template to use for the message.
	//
	// This object isn't supported for campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpoint-campaign-templateconfiguration.html#cfn-pinpoint-campaign-templateconfiguration-voicetemplate
	//
	VoiceTemplate interface{} `field:"optional" json:"voiceTemplate" yaml:"voiceTemplate"`
}


package awspinpoint


// Specifies the message template to use for the message, for each type of channel.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   templateConfigurationProperty := &templateConfigurationProperty{
//   	emailTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	pushTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	smsTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   	voiceTemplate: &templateProperty{
//   		name: jsii.String("name"),
//   		version: jsii.String("version"),
//   	},
//   }
//
type CfnCampaign_TemplateConfigurationProperty struct {
	// The email template to use for the message.
	EmailTemplate interface{} `field:"optional" json:"emailTemplate" yaml:"emailTemplate"`
	// The push notification template to use for the message.
	PushTemplate interface{} `field:"optional" json:"pushTemplate" yaml:"pushTemplate"`
	// The SMS template to use for the message.
	SmsTemplate interface{} `field:"optional" json:"smsTemplate" yaml:"smsTemplate"`
	// The voice template to use for the message.
	//
	// This object isn't supported for campaigns.
	VoiceTemplate interface{} `field:"optional" json:"voiceTemplate" yaml:"voiceTemplate"`
}


package awspinpoint


// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
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
type CfnCampaign_TemplateConfigurationProperty struct {
	// `CfnCampaign.TemplateConfigurationProperty.EmailTemplate`.
	EmailTemplate interface{} `field:"optional" json:"emailTemplate" yaml:"emailTemplate"`
	// `CfnCampaign.TemplateConfigurationProperty.PushTemplate`.
	PushTemplate interface{} `field:"optional" json:"pushTemplate" yaml:"pushTemplate"`
	// `CfnCampaign.TemplateConfigurationProperty.SMSTemplate`.
	SmsTemplate interface{} `field:"optional" json:"smsTemplate" yaml:"smsTemplate"`
	// `CfnCampaign.TemplateConfigurationProperty.VoiceTemplate`.
	VoiceTemplate interface{} `field:"optional" json:"voiceTemplate" yaml:"voiceTemplate"`
}


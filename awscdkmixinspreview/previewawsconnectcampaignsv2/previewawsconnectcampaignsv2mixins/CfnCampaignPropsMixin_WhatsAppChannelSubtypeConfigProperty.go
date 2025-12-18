package previewawsconnectcampaignsv2mixins


// The configuration for the WhatsApp channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   whatsAppChannelSubtypeConfigProperty := &WhatsAppChannelSubtypeConfigProperty{
//   	Capacity: jsii.Number(123),
//   	DefaultOutboundConfig: &WhatsAppOutboundConfigProperty{
//   		ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   		WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   	},
//   	OutboundMode: &WhatsAppOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig.html
//
type CfnCampaignPropsMixin_WhatsAppChannelSubtypeConfigProperty struct {
	// The allocation of WhatsApp capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The default WhatsApp outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"optional" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode for WhatsApp of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-whatsappchannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"optional" json:"outboundMode" yaml:"outboundMode"`
}


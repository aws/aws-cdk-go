package mixinsawsconnectcampaignsv2


// The configuration for the SMS channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   smsChannelSubtypeConfigProperty := &SmsChannelSubtypeConfigProperty{
//   	Capacity: jsii.Number(123),
//   	DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   		ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   		WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   	},
//   	OutboundMode: &SmsOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html
//
type CfnCampaignPropsMixin_SmsChannelSubtypeConfigProperty struct {
	// The allocation of SMS capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The default SMS outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"optional" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode of SMS for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"optional" json:"outboundMode" yaml:"outboundMode"`
}


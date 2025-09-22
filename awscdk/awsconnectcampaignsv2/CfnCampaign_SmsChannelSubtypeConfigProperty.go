package awsconnectcampaignsv2


// The configuration for the SMS channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   smsChannelSubtypeConfigProperty := &SmsChannelSubtypeConfigProperty{
//   	DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   		ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   		WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   	},
//   	OutboundMode: &SmsOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   	},
//
//   	// the properties below are optional
//   	Capacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html
//
type CfnCampaign_SmsChannelSubtypeConfigProperty struct {
	// The default SMS outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"required" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode of SMS for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"required" json:"outboundMode" yaml:"outboundMode"`
	// The allocation of SMS capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}


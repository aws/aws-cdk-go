package awsconnectcampaignsv2


// SMS Channel Subtype config.
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
	// Default SMS outbound config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"required" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// SMS Outbound Mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"required" json:"outboundMode" yaml:"outboundMode"`
	// Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-smschannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-smschannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}


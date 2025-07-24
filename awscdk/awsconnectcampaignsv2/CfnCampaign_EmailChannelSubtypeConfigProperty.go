package awsconnectcampaignsv2


// The configuration for the email channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   emailChannelSubtypeConfigProperty := &EmailChannelSubtypeConfigProperty{
//   	DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   		ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   		WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//
//   		// the properties below are optional
//   		SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   	},
//   	OutboundMode: &EmailOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   	},
//
//   	// the properties below are optional
//   	Capacity: jsii.Number(123),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html
//
type CfnCampaign_EmailChannelSubtypeConfigProperty struct {
	// The default email outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-emailchannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"required" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode for email of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-emailchannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"required" json:"outboundMode" yaml:"outboundMode"`
	// The allocation of email capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-emailchannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
}


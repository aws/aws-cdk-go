package previewawsconnectcampaignsv2mixins


// The configuration for the email channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   emailChannelSubtypeConfigProperty := &EmailChannelSubtypeConfigProperty{
//   	Capacity: jsii.Number(123),
//   	DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   		ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   		SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   		WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   	},
//   	OutboundMode: &EmailOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html
//
type CfnCampaignPropsMixin_EmailChannelSubtypeConfigProperty struct {
	// The allocation of email capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-emailchannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The default email outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-emailchannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"optional" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode for email of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-emailchannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-emailchannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"optional" json:"outboundMode" yaml:"outboundMode"`
}


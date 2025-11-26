package previewawsconnectcampaignsv2mixins


// The configuration for the telephony channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   telephonyChannelSubtypeConfigProperty := &TelephonyChannelSubtypeConfigProperty{
//   	Capacity: jsii.Number(123),
//   	ConnectQueueId: jsii.String("connectQueueId"),
//   	DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   		AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   			AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   			EnableAnswerMachineDetection: jsii.Boolean(false),
//   		},
//   		ConnectContactFlowId: jsii.String("connectContactFlowId"),
//   		ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   		RingTimeout: jsii.Number(123),
//   	},
//   	OutboundMode: &TelephonyOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   		PredictiveConfig: &PredictiveConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   		},
//   		PreviewConfig: &PreviewConfigProperty{
//   			AgentActions: []*string{
//   				jsii.String("agentActions"),
//   			},
//   			BandwidthAllocation: jsii.Number(123),
//   			TimeoutConfig: &TimeoutConfigProperty{
//   				DurationInSeconds: jsii.Number(123),
//   			},
//   		},
//   		ProgressiveConfig: &ProgressiveConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html
//
type CfnCampaignPropsMixin_TelephonyChannelSubtypeConfigProperty struct {
	// The allocation of telephony capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The identifier of the Amazon Connect queue associated with telephony outbound requests of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-connectqueueid
	//
	ConnectQueueId *string `field:"optional" json:"connectQueueId" yaml:"connectQueueId"`
	// The default telephony outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"optional" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode of telephony for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"optional" json:"outboundMode" yaml:"outboundMode"`
}


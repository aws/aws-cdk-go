package awsconnectcampaignsv2


// The configuration for the telephony channel subtype.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   telephonyChannelSubtypeConfigProperty := &TelephonyChannelSubtypeConfigProperty{
//   	DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   		ConnectContactFlowId: jsii.String("connectContactFlowId"),
//
//   		// the properties below are optional
//   		AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   			EnableAnswerMachineDetection: jsii.Boolean(false),
//
//   			// the properties below are optional
//   			AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   		},
//   		ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   		RingTimeout: jsii.Number(123),
//   	},
//   	OutboundMode: &TelephonyOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   		PredictiveConfig: &PredictiveConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   		},
//   		PreviewConfig: &PreviewConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   			TimeoutConfig: &TimeoutConfigProperty{
//   				DurationInSeconds: jsii.Number(123),
//   			},
//
//   			// the properties below are optional
//   			AgentActions: []*string{
//   				jsii.String("agentActions"),
//   			},
//   		},
//   		ProgressiveConfig: &ProgressiveConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
//   		},
//   	},
//
//   	// the properties below are optional
//   	Capacity: jsii.Number(123),
//   	ConnectQueueId: jsii.String("connectQueueId"),
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html
//
type CfnCampaign_TelephonyChannelSubtypeConfigProperty struct {
	// The default telephony outbound configuration of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"required" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// The outbound mode of telephony for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"required" json:"outboundMode" yaml:"outboundMode"`
	// The allocation of telephony capacity between multiple running outbound campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The identifier of the Amazon Connect queue associated with telephony outbound requests of an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-connectqueueid
	//
	ConnectQueueId *string `field:"optional" json:"connectQueueId" yaml:"connectQueueId"`
}


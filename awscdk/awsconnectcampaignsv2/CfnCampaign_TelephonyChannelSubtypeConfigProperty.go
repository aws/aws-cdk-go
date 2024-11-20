package awsconnectcampaignsv2


// Telephony Channel Subtype config.
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
//   	},
//   	OutboundMode: &TelephonyOutboundModeProperty{
//   		AgentlessConfig: agentlessConfig,
//   		PredictiveConfig: &PredictiveConfigProperty{
//   			BandwidthAllocation: jsii.Number(123),
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
	// Default Telephone Outbound config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-defaultoutboundconfig
	//
	DefaultOutboundConfig interface{} `field:"required" json:"defaultOutboundConfig" yaml:"defaultOutboundConfig"`
	// Telephony Outbound Mode.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-outboundmode
	//
	OutboundMode interface{} `field:"required" json:"outboundMode" yaml:"outboundMode"`
	// Allocates outbound capacity for the specific channel of this campaign between multiple active campaigns.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-capacity
	//
	Capacity *float64 `field:"optional" json:"capacity" yaml:"capacity"`
	// The queue for the call.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-telephonychannelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-telephonychannelsubtypeconfig-connectqueueid
	//
	ConnectQueueId *string `field:"optional" json:"connectQueueId" yaml:"connectQueueId"`
}


package awsconnectcampaignsv2


// Contains channel subtype configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   channelSubtypeConfigProperty := &ChannelSubtypeConfigProperty{
//   	Email: &EmailChannelSubtypeConfigProperty{
//   		DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   			ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   			WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//
//   			// the properties below are optional
//   			SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   		},
//   		OutboundMode: &EmailOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   		},
//
//   		// the properties below are optional
//   		Capacity: jsii.Number(123),
//   	},
//   	Sms: &SmsChannelSubtypeConfigProperty{
//   		DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   			ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   			WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   		},
//   		OutboundMode: &SmsOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   		},
//
//   		// the properties below are optional
//   		Capacity: jsii.Number(123),
//   	},
//   	Telephony: &TelephonyChannelSubtypeConfigProperty{
//   		DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   			ConnectContactFlowId: jsii.String("connectContactFlowId"),
//
//   			// the properties below are optional
//   			AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   				EnableAnswerMachineDetection: jsii.Boolean(false),
//
//   				// the properties below are optional
//   				AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   			},
//   			ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   			RingTimeout: jsii.Number(123),
//   		},
//   		OutboundMode: &TelephonyOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   			PredictiveConfig: &PredictiveConfigProperty{
//   				BandwidthAllocation: jsii.Number(123),
//   			},
//   			PreviewConfig: &PreviewConfigProperty{
//   				BandwidthAllocation: jsii.Number(123),
//   				TimeoutConfig: &TimeoutConfigProperty{
//   					DurationInSeconds: jsii.Number(123),
//   				},
//
//   				// the properties below are optional
//   				AgentActions: []*string{
//   					jsii.String("agentActions"),
//   				},
//   			},
//   			ProgressiveConfig: &ProgressiveConfigProperty{
//   				BandwidthAllocation: jsii.Number(123),
//   			},
//   		},
//
//   		// the properties below are optional
//   		Capacity: jsii.Number(123),
//   		ConnectQueueId: jsii.String("connectQueueId"),
//   	},
//   	WhatsApp: &WhatsAppChannelSubtypeConfigProperty{
//   		DefaultOutboundConfig: &WhatsAppOutboundConfigProperty{
//   			ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   			WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   		},
//   		OutboundMode: &WhatsAppOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   		},
//
//   		// the properties below are optional
//   		Capacity: jsii.Number(123),
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.html
//
type CfnCampaign_ChannelSubtypeConfigProperty struct {
	// The configuration of the email channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-channelsubtypeconfig-email
	//
	Email interface{} `field:"optional" json:"email" yaml:"email"`
	// The configuration of the SMS channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-channelsubtypeconfig-sms
	//
	Sms interface{} `field:"optional" json:"sms" yaml:"sms"`
	// The configuration of the telephony channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-channelsubtypeconfig-telephony
	//
	Telephony interface{} `field:"optional" json:"telephony" yaml:"telephony"`
	// The configuration of the WhatsApp channel subtype.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.html#cfn-connectcampaignsv2-campaign-channelsubtypeconfig-whatsapp
	//
	WhatsApp interface{} `field:"optional" json:"whatsApp" yaml:"whatsApp"`
}


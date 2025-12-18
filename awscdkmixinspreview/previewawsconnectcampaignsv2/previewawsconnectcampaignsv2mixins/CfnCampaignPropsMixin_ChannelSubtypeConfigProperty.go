package previewawsconnectcampaignsv2mixins


// Contains channel subtype configuration for an outbound campaign.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   channelSubtypeConfigProperty := &ChannelSubtypeConfigProperty{
//   	Email: &EmailChannelSubtypeConfigProperty{
//   		Capacity: jsii.Number(123),
//   		DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   			ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   			SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   			WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   		},
//   		OutboundMode: &EmailOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   		},
//   	},
//   	Sms: &SmsChannelSubtypeConfigProperty{
//   		Capacity: jsii.Number(123),
//   		DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   			ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   			WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   		},
//   		OutboundMode: &SmsOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   		},
//   	},
//   	Telephony: &TelephonyChannelSubtypeConfigProperty{
//   		Capacity: jsii.Number(123),
//   		ConnectQueueId: jsii.String("connectQueueId"),
//   		DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   			AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   				AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   				EnableAnswerMachineDetection: jsii.Boolean(false),
//   			},
//   			ConnectContactFlowId: jsii.String("connectContactFlowId"),
//   			ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   			RingTimeout: jsii.Number(123),
//   		},
//   		OutboundMode: &TelephonyOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   			PredictiveConfig: &PredictiveConfigProperty{
//   				BandwidthAllocation: jsii.Number(123),
//   			},
//   			PreviewConfig: &PreviewConfigProperty{
//   				AgentActions: []*string{
//   					jsii.String("agentActions"),
//   				},
//   				BandwidthAllocation: jsii.Number(123),
//   				TimeoutConfig: &TimeoutConfigProperty{
//   					DurationInSeconds: jsii.Number(123),
//   				},
//   			},
//   			ProgressiveConfig: &ProgressiveConfigProperty{
//   				BandwidthAllocation: jsii.Number(123),
//   			},
//   		},
//   	},
//   	WhatsApp: &WhatsAppChannelSubtypeConfigProperty{
//   		Capacity: jsii.Number(123),
//   		DefaultOutboundConfig: &WhatsAppOutboundConfigProperty{
//   			ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   			WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   		},
//   		OutboundMode: &WhatsAppOutboundModeProperty{
//   			AgentlessConfig: agentlessConfig,
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-connectcampaignsv2-campaign-channelsubtypeconfig.html
//
type CfnCampaignPropsMixin_ChannelSubtypeConfigProperty struct {
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


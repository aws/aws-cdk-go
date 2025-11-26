package previewawsconnectcampaignsv2mixins

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for CfnCampaignPropsMixin.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   cfnCampaignMixinProps := &CfnCampaignMixinProps{
//   	ChannelSubtypeConfig: &ChannelSubtypeConfigProperty{
//   		Email: &EmailChannelSubtypeConfigProperty{
//   			Capacity: jsii.Number(123),
//   			DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   				ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   				SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   			},
//   			OutboundMode: &EmailOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   			},
//   		},
//   		Sms: &SmsChannelSubtypeConfigProperty{
//   			Capacity: jsii.Number(123),
//   			DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   				ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   			},
//   			OutboundMode: &SmsOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   			},
//   		},
//   		Telephony: &TelephonyChannelSubtypeConfigProperty{
//   			Capacity: jsii.Number(123),
//   			ConnectQueueId: jsii.String("connectQueueId"),
//   			DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   				AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   					AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   					EnableAnswerMachineDetection: jsii.Boolean(false),
//   				},
//   				ConnectContactFlowId: jsii.String("connectContactFlowId"),
//   				ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   				RingTimeout: jsii.Number(123),
//   			},
//   			OutboundMode: &TelephonyOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   				PredictiveConfig: &PredictiveConfigProperty{
//   					BandwidthAllocation: jsii.Number(123),
//   				},
//   				PreviewConfig: &PreviewConfigProperty{
//   					AgentActions: []*string{
//   						jsii.String("agentActions"),
//   					},
//   					BandwidthAllocation: jsii.Number(123),
//   					TimeoutConfig: &TimeoutConfigProperty{
//   						DurationInSeconds: jsii.Number(123),
//   					},
//   				},
//   				ProgressiveConfig: &ProgressiveConfigProperty{
//   					BandwidthAllocation: jsii.Number(123),
//   				},
//   			},
//   		},
//   	},
//   	CommunicationLimitsOverride: &CommunicationLimitsConfigProperty{
//   		AllChannelsSubtypes: &CommunicationLimitsProperty{
//   			CommunicationLimitList: []interface{}{
//   				&CommunicationLimitProperty{
//   					Frequency: jsii.Number(123),
//   					MaxCountPerRecipient: jsii.Number(123),
//   					Unit: jsii.String("unit"),
//   				},
//   			},
//   		},
//   		InstanceLimitsHandling: jsii.String("instanceLimitsHandling"),
//   	},
//   	CommunicationTimeConfig: &CommunicationTimeConfigProperty{
//   		Email: &TimeWindowProperty{
//   			OpenHours: &OpenHoursProperty{
//   				DailyHours: []interface{}{
//   					&DailyHourProperty{
//   						Key: jsii.String("key"),
//   						Value: []interface{}{
//   							&TimeRangeProperty{
//   								EndTime: jsii.String("endTime"),
//   								StartTime: jsii.String("startTime"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						Name: jsii.String("name"),
//   						StartDate: jsii.String("startDate"),
//   					},
//   				},
//   			},
//   		},
//   		LocalTimeZoneConfig: &LocalTimeZoneConfigProperty{
//   			DefaultTimeZone: jsii.String("defaultTimeZone"),
//   			LocalTimeZoneDetection: []*string{
//   				jsii.String("localTimeZoneDetection"),
//   			},
//   		},
//   		Sms: &TimeWindowProperty{
//   			OpenHours: &OpenHoursProperty{
//   				DailyHours: []interface{}{
//   					&DailyHourProperty{
//   						Key: jsii.String("key"),
//   						Value: []interface{}{
//   							&TimeRangeProperty{
//   								EndTime: jsii.String("endTime"),
//   								StartTime: jsii.String("startTime"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						Name: jsii.String("name"),
//   						StartDate: jsii.String("startDate"),
//   					},
//   				},
//   			},
//   		},
//   		Telephony: &TimeWindowProperty{
//   			OpenHours: &OpenHoursProperty{
//   				DailyHours: []interface{}{
//   					&DailyHourProperty{
//   						Key: jsii.String("key"),
//   						Value: []interface{}{
//   							&TimeRangeProperty{
//   								EndTime: jsii.String("endTime"),
//   								StartTime: jsii.String("startTime"),
//   							},
//   						},
//   					},
//   				},
//   			},
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						Name: jsii.String("name"),
//   						StartDate: jsii.String("startDate"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ConnectCampaignFlowArn: jsii.String("connectCampaignFlowArn"),
//   	ConnectInstanceId: jsii.String("connectInstanceId"),
//   	Name: jsii.String("name"),
//   	Schedule: &ScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		RefreshFrequency: jsii.String("refreshFrequency"),
//   		StartTime: jsii.String("startTime"),
//   	},
//   	Source: &SourceProperty{
//   		CustomerProfilesSegmentArn: jsii.String("customerProfilesSegmentArn"),
//   		EventTrigger: &EventTriggerProperty{
//   			CustomerProfilesDomainArn: jsii.String("customerProfilesDomainArn"),
//   		},
//   	},
//   	Tags: []CfnTag{
//   		&CfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html
//
type CfnCampaignMixinProps struct {
	// Contains channel subtype configuration for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-channelsubtypeconfig
	//
	ChannelSubtypeConfig interface{} `field:"optional" json:"channelSubtypeConfig" yaml:"channelSubtypeConfig"`
	// Communication limits configuration for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-communicationlimitsoverride
	//
	CommunicationLimitsOverride interface{} `field:"optional" json:"communicationLimitsOverride" yaml:"communicationLimitsOverride"`
	// Contains communication time configuration for an outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig
	//
	CommunicationTimeConfig interface{} `field:"optional" json:"communicationTimeConfig" yaml:"communicationTimeConfig"`
	// The Amazon Resource Name (ARN) of the Amazon Connect campaign flow associated with the outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-connectcampaignflowarn
	//
	ConnectCampaignFlowArn *string `field:"optional" json:"connectCampaignFlowArn" yaml:"connectCampaignFlowArn"`
	// The identifier of the Amazon Connect instance.
	//
	// You can find the `instanceId` in the ARN of the instance.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-connectinstanceid
	//
	ConnectInstanceId *string `field:"optional" json:"connectInstanceId" yaml:"connectInstanceId"`
	// The name of the outbound campaign.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-name
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Contains the schedule configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// Contains source configuration.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// The tags used to organize, track, or control access for this resource.
	//
	// For example, `{ "tags": {"key1":"value1", "key2":"value2"} }` .
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


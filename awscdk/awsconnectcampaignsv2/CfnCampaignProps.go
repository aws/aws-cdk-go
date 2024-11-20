package awsconnectcampaignsv2

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
)

// Properties for defining a `CfnCampaign`.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var agentlessConfig interface{}
//
//   cfnCampaignProps := &CfnCampaignProps{
//   	ChannelSubtypeConfig: &ChannelSubtypeConfigProperty{
//   		Email: &EmailChannelSubtypeConfigProperty{
//   			DefaultOutboundConfig: &EmailOutboundConfigProperty{
//   				ConnectSourceEmailAddress: jsii.String("connectSourceEmailAddress"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//
//   				// the properties below are optional
//   				SourceEmailAddressDisplayName: jsii.String("sourceEmailAddressDisplayName"),
//   			},
//   			OutboundMode: &EmailOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   			},
//
//   			// the properties below are optional
//   			Capacity: jsii.Number(123),
//   		},
//   		Sms: &SmsChannelSubtypeConfigProperty{
//   			DefaultOutboundConfig: &SmsOutboundConfigProperty{
//   				ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   			},
//   			OutboundMode: &SmsOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   			},
//
//   			// the properties below are optional
//   			Capacity: jsii.Number(123),
//   		},
//   		Telephony: &TelephonyChannelSubtypeConfigProperty{
//   			DefaultOutboundConfig: &TelephonyOutboundConfigProperty{
//   				ConnectContactFlowId: jsii.String("connectContactFlowId"),
//
//   				// the properties below are optional
//   				AnswerMachineDetectionConfig: &AnswerMachineDetectionConfigProperty{
//   					EnableAnswerMachineDetection: jsii.Boolean(false),
//
//   					// the properties below are optional
//   					AwaitAnswerMachinePrompt: jsii.Boolean(false),
//   				},
//   				ConnectSourcePhoneNumber: jsii.String("connectSourcePhoneNumber"),
//   			},
//   			OutboundMode: &TelephonyOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
//   				PredictiveConfig: &PredictiveConfigProperty{
//   					BandwidthAllocation: jsii.Number(123),
//   				},
//   				ProgressiveConfig: &ProgressiveConfigProperty{
//   					BandwidthAllocation: jsii.Number(123),
//   				},
//   			},
//
//   			// the properties below are optional
//   			Capacity: jsii.Number(123),
//   			ConnectQueueId: jsii.String("connectQueueId"),
//   		},
//   	},
//   	ConnectInstanceId: jsii.String("connectInstanceId"),
//   	Name: jsii.String("name"),
//
//   	// the properties below are optional
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
//   	},
//   	CommunicationTimeConfig: &CommunicationTimeConfigProperty{
//   		LocalTimeZoneConfig: &LocalTimeZoneConfigProperty{
//   			DefaultTimeZone: jsii.String("defaultTimeZone"),
//   			LocalTimeZoneDetection: []*string{
//   				jsii.String("localTimeZoneDetection"),
//   			},
//   		},
//
//   		// the properties below are optional
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
//
//   			// the properties below are optional
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						StartDate: jsii.String("startDate"),
//
//   						// the properties below are optional
//   						Name: jsii.String("name"),
//   					},
//   				},
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
//
//   			// the properties below are optional
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						StartDate: jsii.String("startDate"),
//
//   						// the properties below are optional
//   						Name: jsii.String("name"),
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
//
//   			// the properties below are optional
//   			RestrictedPeriods: &RestrictedPeriodsProperty{
//   				RestrictedPeriodList: []interface{}{
//   					&RestrictedPeriodProperty{
//   						EndDate: jsii.String("endDate"),
//   						StartDate: jsii.String("startDate"),
//
//   						// the properties below are optional
//   						Name: jsii.String("name"),
//   					},
//   				},
//   			},
//   		},
//   	},
//   	ConnectCampaignFlowArn: jsii.String("connectCampaignFlowArn"),
//   	Schedule: &ScheduleProperty{
//   		EndTime: jsii.String("endTime"),
//   		StartTime: jsii.String("startTime"),
//
//   		// the properties below are optional
//   		RefreshFrequency: jsii.String("refreshFrequency"),
//   	},
//   	Source: &SourceProperty{
//   		CustomerProfilesSegmentArn: jsii.String("customerProfilesSegmentArn"),
//   	},
//   	Tags: []cfnTag{
//   		&cfnTag{
//   			Key: jsii.String("key"),
//   			Value: jsii.String("value"),
//   		},
//   	},
//   }
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html
//
type CfnCampaignProps struct {
	// The possible types of channel subtype config parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-channelsubtypeconfig
	//
	ChannelSubtypeConfig interface{} `field:"required" json:"channelSubtypeConfig" yaml:"channelSubtypeConfig"`
	// Amazon Connect Instance Id.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-connectinstanceid
	//
	ConnectInstanceId *string `field:"required" json:"connectInstanceId" yaml:"connectInstanceId"`
	// Campaign name.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-name
	//
	Name *string `field:"required" json:"name" yaml:"name"`
	// Communication limits config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-communicationlimitsoverride
	//
	CommunicationLimitsOverride interface{} `field:"optional" json:"communicationLimitsOverride" yaml:"communicationLimitsOverride"`
	// Campaign communication time config.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-communicationtimeconfig
	//
	CommunicationTimeConfig interface{} `field:"optional" json:"communicationTimeConfig" yaml:"communicationTimeConfig"`
	// Arn.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-connectcampaignflowarn
	//
	ConnectCampaignFlowArn *string `field:"optional" json:"connectCampaignFlowArn" yaml:"connectCampaignFlowArn"`
	// Campaign schedule.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-schedule
	//
	Schedule interface{} `field:"optional" json:"schedule" yaml:"schedule"`
	// The possible types of channel config parameters.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-source
	//
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// One or more tags.
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html#cfn-connectcampaignsv2-campaign-tags
	//
	Tags *[]*awscdk.CfnTag `field:"optional" json:"tags" yaml:"tags"`
}


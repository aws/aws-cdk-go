package previewawsconnectcampaignsv2mixins

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/core"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/mixins"
	"github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/previewawsconnectcampaignsv2/previewawsconnectcampaignsv2mixins/internal"
	"github.com/aws/constructs-go/constructs/v10"
)

// Creates an outbound campaign.
//
// > - For users to be able to view or edit a campaign at a later date by using the Amazon Connect user interface, you must add the instance ID as a tag. For example, `{ "tags": {"owner": "arn:aws:connect:{REGION}:{AWS_ACCOUNT_ID}:instance/{CONNECT_INSTANCE_ID}"}}` .
// > - After a campaign is created, you can't add/remove source.
// > - Configuring maximum ring time is not supported for the Preview dial mode.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   var agentlessConfig interface{}
//
//   cfnCampaignPropsMixin := awscdkmixinspreview.Mixins.NewCfnCampaignPropsMixin(&CfnCampaignMixinProps{
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
//   		WhatsApp: &WhatsAppChannelSubtypeConfigProperty{
//   			Capacity: jsii.Number(123),
//   			DefaultOutboundConfig: &WhatsAppOutboundConfigProperty{
//   				ConnectSourcePhoneNumberArn: jsii.String("connectSourcePhoneNumberArn"),
//   				WisdomTemplateArn: jsii.String("wisdomTemplateArn"),
//   			},
//   			OutboundMode: &WhatsAppOutboundModeProperty{
//   				AgentlessConfig: agentlessConfig,
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
//   		WhatsApp: &TimeWindowProperty{
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
//   	Type: jsii.String("type"),
//   }, &CfnPropertyMixinOptions{
//   	Strategy: awscdkmixinspreview.Mixins.PropertyMergeStrategy_OVERRIDE,
//   })
//
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-connectcampaignsv2-campaign.html
//
type CfnCampaignPropsMixin interface {
	core.Mixin
	core.IMixin
	Props() *CfnCampaignMixinProps
	Strategy() mixins.PropertyMergeStrategy
	// Apply the mixin properties to the construct.
	ApplyTo(construct constructs.IConstruct) constructs.IConstruct
	// Check if this mixin supports the given construct.
	Supports(construct constructs.IConstruct) *bool
}

// The jsii proxy struct for CfnCampaignPropsMixin
type jsiiProxy_CfnCampaignPropsMixin struct {
	internal.Type__coreMixin
	internal.Type__coreIMixin
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Props() *CfnCampaignMixinProps {
	var returns *CfnCampaignMixinProps
	_jsii_.Get(
		j,
		"props",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CfnCampaignPropsMixin) Strategy() mixins.PropertyMergeStrategy {
	var returns mixins.PropertyMergeStrategy
	_jsii_.Get(
		j,
		"strategy",
		&returns,
	)
	return returns
}


// Create a mixin to apply properties to `AWS::ConnectCampaignsV2::Campaign`.
func NewCfnCampaignPropsMixin(props *CfnCampaignMixinProps, options *mixins.CfnPropertyMixinOptions) CfnCampaignPropsMixin {
	_init_.Initialize()

	if err := validateNewCfnCampaignPropsMixinParameters(props, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_CfnCampaignPropsMixin{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		&j,
	)

	return &j
}

// Create a mixin to apply properties to `AWS::ConnectCampaignsV2::Campaign`.
func NewCfnCampaignPropsMixin_Override(c CfnCampaignPropsMixin, props *CfnCampaignMixinProps, options *mixins.CfnPropertyMixinOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin",
		[]interface{}{props, options},
		c,
	)
}

// Checks if `x` is a Mixin.
//
// Returns: true if `x` is an object created from a class which extends `Mixin`.
// Experimental.
func CfnCampaignPropsMixin_IsMixin(x interface{}) *bool {
	_init_.Initialize()

	if err := validateCfnCampaignPropsMixin_IsMixinParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin",
		"isMixin",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func CfnCampaignPropsMixin_CFN_PROPERTY_KEYS() *[]*string {
	_init_.Initialize()
	var returns *[]*string
	_jsii_.StaticGet(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin",
		"CFN_PROPERTY_KEYS",
		&returns,
	)
	return returns
}

func (c *jsiiProxy_CfnCampaignPropsMixin) ApplyTo(construct constructs.IConstruct) constructs.IConstruct {
	if err := c.validateApplyToParameters(construct); err != nil {
		panic(err)
	}
	var returns constructs.IConstruct

	_jsii_.Invoke(
		c,
		"applyTo",
		[]interface{}{construct},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CfnCampaignPropsMixin) Supports(construct constructs.IConstruct) *bool {
	if err := c.validateSupportsParameters(construct); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.Invoke(
		c,
		"supports",
		[]interface{}{construct},
		&returns,
	)

	return returns
}


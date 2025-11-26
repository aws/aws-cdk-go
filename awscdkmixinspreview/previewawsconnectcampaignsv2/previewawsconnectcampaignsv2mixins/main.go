package previewawsconnectcampaignsv2mixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignMixinProps",
		reflect.TypeOf((*CfnCampaignMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin",
		reflect.TypeOf((*CfnCampaignPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCampaignPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.AnswerMachineDetectionConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_AnswerMachineDetectionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.ChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.CommunicationLimitProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_CommunicationLimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.CommunicationLimitsConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_CommunicationLimitsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.CommunicationLimitsProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_CommunicationLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.CommunicationTimeConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_CommunicationTimeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.DailyHourProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_DailyHourProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.EmailChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_EmailChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.EmailOutboundConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_EmailOutboundConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.EmailOutboundModeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_EmailOutboundModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.EventTriggerProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_EventTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.LocalTimeZoneConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_LocalTimeZoneConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.OpenHoursProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_OpenHoursProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.PredictiveConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_PredictiveConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.PreviewConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_PreviewConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.ProgressiveConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ProgressiveConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.RestrictedPeriodProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_RestrictedPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.RestrictedPeriodsProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_RestrictedPeriodsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.SmsChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SmsChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.SmsOutboundConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SmsOutboundConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.SmsOutboundModeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SmsOutboundModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.TelephonyChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TelephonyChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.TelephonyOutboundConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TelephonyOutboundConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.TelephonyOutboundModeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TelephonyOutboundModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.TimeRangeProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.TimeWindowProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_connectcampaignsv2.mixins.CfnCampaignPropsMixin.TimeoutConfigProperty",
		reflect.TypeOf((*CfnCampaignPropsMixin_TimeoutConfigProperty)(nil)).Elem(),
	)
}

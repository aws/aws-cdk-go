package awsconnectcampaignsv2

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CampaignReference",
		reflect.TypeOf((*CampaignReference)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign",
		reflect.TypeOf((*CfnCampaign)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addDeletionOverride", GoMethod: "AddDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addDependency", GoMethod: "AddDependency"},
			_jsii_.MemberMethod{JsiiMethod: "addDependsOn", GoMethod: "AddDependsOn"},
			_jsii_.MemberMethod{JsiiMethod: "addMetadata", GoMethod: "AddMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyDeletionOverride", GoMethod: "AddPropertyDeletionOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addPropertyOverride", GoMethod: "AddPropertyOverride"},
			_jsii_.MemberMethod{JsiiMethod: "applyRemovalPolicy", GoMethod: "ApplyRemovalPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "attrArn", GoGetter: "AttrArn"},
			_jsii_.MemberProperty{JsiiProperty: "campaignRef", GoGetter: "CampaignRef"},
			_jsii_.MemberProperty{JsiiProperty: "cdkTagManager", GoGetter: "CdkTagManager"},
			_jsii_.MemberProperty{JsiiProperty: "cfnOptions", GoGetter: "CfnOptions"},
			_jsii_.MemberProperty{JsiiProperty: "cfnProperties", GoGetter: "CfnProperties"},
			_jsii_.MemberProperty{JsiiProperty: "cfnResourceType", GoGetter: "CfnResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "channelSubtypeConfig", GoGetter: "ChannelSubtypeConfig"},
			_jsii_.MemberProperty{JsiiProperty: "communicationLimitsOverride", GoGetter: "CommunicationLimitsOverride"},
			_jsii_.MemberProperty{JsiiProperty: "communicationTimeConfig", GoGetter: "CommunicationTimeConfig"},
			_jsii_.MemberProperty{JsiiProperty: "connectCampaignFlowArn", GoGetter: "ConnectCampaignFlowArn"},
			_jsii_.MemberProperty{JsiiProperty: "connectInstanceId", GoGetter: "ConnectInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberMethod{JsiiMethod: "getAtt", GoMethod: "GetAtt"},
			_jsii_.MemberMethod{JsiiMethod: "getMetadata", GoMethod: "GetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "inspect", GoMethod: "Inspect"},
			_jsii_.MemberProperty{JsiiProperty: "logicalId", GoGetter: "LogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "obtainDependencies", GoMethod: "ObtainDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "obtainResourceDependencies", GoMethod: "ObtainResourceDependencies"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "ref", GoGetter: "Ref"},
			_jsii_.MemberMethod{JsiiMethod: "removeDependency", GoMethod: "RemoveDependency"},
			_jsii_.MemberMethod{JsiiMethod: "renderProperties", GoMethod: "RenderProperties"},
			_jsii_.MemberMethod{JsiiMethod: "replaceDependency", GoMethod: "ReplaceDependency"},
			_jsii_.MemberProperty{JsiiProperty: "schedule", GoGetter: "Schedule"},
			_jsii_.MemberMethod{JsiiMethod: "shouldSynthesize", GoMethod: "ShouldSynthesize"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "stack", GoGetter: "Stack"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperites", GoGetter: "UpdatedProperites"},
			_jsii_.MemberProperty{JsiiProperty: "updatedProperties", GoGetter: "UpdatedProperties"},
			_jsii_.MemberMethod{JsiiMethod: "validateProperties", GoMethod: "ValidateProperties"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCampaign{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkCfnResource)
			_jsii_.InitJsiiProxy(&j.jsiiProxy_ICampaignRef)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIInspectable)
			_jsii_.InitJsiiProxy(&j.Type__awscdkITaggableV2)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.AnswerMachineDetectionConfigProperty",
		reflect.TypeOf((*CfnCampaign_AnswerMachineDetectionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.ChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaign_ChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.CommunicationLimitProperty",
		reflect.TypeOf((*CfnCampaign_CommunicationLimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.CommunicationLimitsConfigProperty",
		reflect.TypeOf((*CfnCampaign_CommunicationLimitsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.CommunicationLimitsProperty",
		reflect.TypeOf((*CfnCampaign_CommunicationLimitsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.CommunicationTimeConfigProperty",
		reflect.TypeOf((*CfnCampaign_CommunicationTimeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.DailyHourProperty",
		reflect.TypeOf((*CfnCampaign_DailyHourProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.EmailChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaign_EmailChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.EmailOutboundConfigProperty",
		reflect.TypeOf((*CfnCampaign_EmailOutboundConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.EmailOutboundModeProperty",
		reflect.TypeOf((*CfnCampaign_EmailOutboundModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.EventTriggerProperty",
		reflect.TypeOf((*CfnCampaign_EventTriggerProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.LocalTimeZoneConfigProperty",
		reflect.TypeOf((*CfnCampaign_LocalTimeZoneConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.OpenHoursProperty",
		reflect.TypeOf((*CfnCampaign_OpenHoursProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.PredictiveConfigProperty",
		reflect.TypeOf((*CfnCampaign_PredictiveConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.PreviewConfigProperty",
		reflect.TypeOf((*CfnCampaign_PreviewConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.ProgressiveConfigProperty",
		reflect.TypeOf((*CfnCampaign_ProgressiveConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.RestrictedPeriodProperty",
		reflect.TypeOf((*CfnCampaign_RestrictedPeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.RestrictedPeriodsProperty",
		reflect.TypeOf((*CfnCampaign_RestrictedPeriodsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.ScheduleProperty",
		reflect.TypeOf((*CfnCampaign_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.SmsChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaign_SmsChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.SmsOutboundConfigProperty",
		reflect.TypeOf((*CfnCampaign_SmsOutboundConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.SmsOutboundModeProperty",
		reflect.TypeOf((*CfnCampaign_SmsOutboundModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.SourceProperty",
		reflect.TypeOf((*CfnCampaign_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.TelephonyChannelSubtypeConfigProperty",
		reflect.TypeOf((*CfnCampaign_TelephonyChannelSubtypeConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.TelephonyOutboundConfigProperty",
		reflect.TypeOf((*CfnCampaign_TelephonyOutboundConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.TelephonyOutboundModeProperty",
		reflect.TypeOf((*CfnCampaign_TelephonyOutboundModeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.TimeRangeProperty",
		reflect.TypeOf((*CfnCampaign_TimeRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.TimeWindowProperty",
		reflect.TypeOf((*CfnCampaign_TimeWindowProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaign.TimeoutConfigProperty",
		reflect.TypeOf((*CfnCampaign_TimeoutConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"aws-cdk-lib.aws_connectcampaignsv2.CfnCampaignProps",
		reflect.TypeOf((*CfnCampaignProps)(nil)).Elem(),
	)
	_jsii_.RegisterInterface(
		"aws-cdk-lib.aws_connectcampaignsv2.ICampaignRef",
		reflect.TypeOf((*ICampaignRef)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "campaignRef", GoGetter: "CampaignRef"},
			_jsii_.MemberProperty{JsiiProperty: "env", GoGetter: "Env"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
		},
		func() interface{} {
			j := jsiiProxy_ICampaignRef{}
			_jsii_.InitJsiiProxy(&j.Type__constructsIConstruct)
			_jsii_.InitJsiiProxy(&j.Type__awscdkIEnvironmentAware)
			return &j
		},
	)
}

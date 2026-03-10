package previewawsiotanalyticsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.ChannelStorage",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ChannelStorage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.CustomerManagedS3",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_CustomerManagedS3)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.LoggingOptions",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_LoggingOptions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.MathType",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_MathType)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.PipelineActivity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_PipelineActivity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.QueryAction",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_QueryAction)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.RequestParametersItem1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.RetentionPeriod",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RetentionPeriod)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.RetentionPeriod1",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RetentionPeriod1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.Settings",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Settings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.Timeseries",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Timeseries)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.TimeseriesItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_TimeseriesItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.AWSAPICallViaCloudTrail.WebIdFederationData",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_WebIdFederationData)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.DatasetEvents",
		reflect.TypeOf((*DatasetEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "ioTAnalyticsDataSetLifeCycleNotificationPattern", GoMethod: "IoTAnalyticsDataSetLifeCycleNotificationPattern"},
		},
		func() interface{} {
			return &jsiiProxy_DatasetEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.IoTAnalyticsDataSetLifeCycleNotification",
		reflect.TypeOf((*IoTAnalyticsDataSetLifeCycleNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_IoTAnalyticsDataSetLifeCycleNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.IoTAnalyticsDataSetLifeCycleNotification.IoTAnalyticsDataSetLifeCycleNotificationProps",
		reflect.TypeOf((*IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps)(nil)).Elem(),
	)
}

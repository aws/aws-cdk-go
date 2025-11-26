package previewawsiotanalyticsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
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
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.DatasetEvents.IoTAnalyticsDataSetLifeCycleNotification",
		reflect.TypeOf((*DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotanalytics.events.DatasetEvents.IoTAnalyticsDataSetLifeCycleNotification.IoTAnalyticsDataSetLifeCycleNotificationProps",
		reflect.TypeOf((*DatasetEvents_IoTAnalyticsDataSetLifeCycleNotification_IoTAnalyticsDataSetLifeCycleNotificationProps)(nil)).Elem(),
	)
}

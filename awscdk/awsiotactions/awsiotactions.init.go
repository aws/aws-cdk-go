package awsiotactions

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.CloudWatchLogsAction",
		reflect.TypeOf((*CloudWatchLogsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchLogsAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.CloudWatchLogsActionProps",
		reflect.TypeOf((*CloudWatchLogsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.CloudWatchPutMetricAction",
		reflect.TypeOf((*CloudWatchPutMetricAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchPutMetricAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.CloudWatchPutMetricActionProps",
		reflect.TypeOf((*CloudWatchPutMetricActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.CloudWatchSetAlarmStateAction",
		reflect.TypeOf((*CloudWatchSetAlarmStateAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchSetAlarmStateAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.CloudWatchSetAlarmStateActionProps",
		reflect.TypeOf((*CloudWatchSetAlarmStateActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.CommonActionProps",
		reflect.TypeOf((*CommonActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.FirehosePutRecordAction",
		reflect.TypeOf((*FirehosePutRecordAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FirehosePutRecordAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.FirehosePutRecordActionProps",
		reflect.TypeOf((*FirehosePutRecordActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_iot_actions.FirehoseRecordSeparator",
		reflect.TypeOf((*FirehoseRecordSeparator)(nil)).Elem(),
		map[string]interface{}{
			"NEWLINE": FirehoseRecordSeparator_NEWLINE,
			"TAB": FirehoseRecordSeparator_TAB,
			"WINDOWS_NEWLINE": FirehoseRecordSeparator_WINDOWS_NEWLINE,
			"COMMA": FirehoseRecordSeparator_COMMA,
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.KinesisPutRecordAction",
		reflect.TypeOf((*KinesisPutRecordAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisPutRecordAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.KinesisPutRecordActionProps",
		reflect.TypeOf((*KinesisPutRecordActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.LambdaFunctionAction",
		reflect.TypeOf((*LambdaFunctionAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunctionAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.S3PutObjectAction",
		reflect.TypeOf((*S3PutObjectAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3PutObjectAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.S3PutObjectActionProps",
		reflect.TypeOf((*S3PutObjectActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.SqsQueueAction",
		reflect.TypeOf((*SqsQueueAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SqsQueueAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.SqsQueueActionProps",
		reflect.TypeOf((*SqsQueueActionProps)(nil)).Elem(),
	)
}

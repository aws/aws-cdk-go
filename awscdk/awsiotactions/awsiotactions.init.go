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
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.CommonActionProps",
		reflect.TypeOf((*CommonActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"monocdk.aws_iot_actions.FirehoseStreamAction",
		reflect.TypeOf((*FirehoseStreamAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FirehoseStreamAction{}
			_jsii_.InitJsiiProxy(&j.Type__awsiotIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"monocdk.aws_iot_actions.FirehoseStreamActionProps",
		reflect.TypeOf((*FirehoseStreamActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"monocdk.aws_iot_actions.FirehoseStreamRecordSeparator",
		reflect.TypeOf((*FirehoseStreamRecordSeparator)(nil)).Elem(),
		map[string]interface{}{
			"NEWLINE": FirehoseStreamRecordSeparator_NEWLINE,
			"TAB": FirehoseStreamRecordSeparator_TAB,
			"WINDOWS_NEWLINE": FirehoseStreamRecordSeparator_WINDOWS_NEWLINE,
			"COMMA": FirehoseStreamRecordSeparator_COMMA,
		},
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
}

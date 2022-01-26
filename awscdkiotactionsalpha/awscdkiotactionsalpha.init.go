package awscdkiotactionsalpha

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchLogsAction",
		reflect.TypeOf((*CloudWatchLogsAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchLogsAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchLogsActionProps",
		reflect.TypeOf((*CloudWatchLogsActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchPutMetricAction",
		reflect.TypeOf((*CloudWatchPutMetricAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchPutMetricAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchPutMetricActionProps",
		reflect.TypeOf((*CloudWatchPutMetricActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchSetAlarmStateAction",
		reflect.TypeOf((*CloudWatchSetAlarmStateAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_CloudWatchSetAlarmStateAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.CloudWatchSetAlarmStateActionProps",
		reflect.TypeOf((*CloudWatchSetAlarmStateActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.CommonActionProps",
		reflect.TypeOf((*CommonActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.FirehosePutRecordAction",
		reflect.TypeOf((*FirehosePutRecordAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FirehosePutRecordAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.FirehosePutRecordActionProps",
		reflect.TypeOf((*FirehosePutRecordActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-iot-actions-alpha.FirehoseRecordSeparator",
		reflect.TypeOf((*FirehoseRecordSeparator)(nil)).Elem(),
		map[string]interface{}{
			"NEWLINE": FirehoseRecordSeparator_NEWLINE,
			"TAB": FirehoseRecordSeparator_TAB,
			"WINDOWS_NEWLINE": FirehoseRecordSeparator_WINDOWS_NEWLINE,
			"COMMA": FirehoseRecordSeparator_COMMA,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.KinesisPutRecordAction",
		reflect.TypeOf((*KinesisPutRecordAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_KinesisPutRecordAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.KinesisPutRecordActionProps",
		reflect.TypeOf((*KinesisPutRecordActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.LambdaFunctionAction",
		reflect.TypeOf((*LambdaFunctionAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_LambdaFunctionAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.S3PutObjectAction",
		reflect.TypeOf((*S3PutObjectAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_S3PutObjectAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.S3PutObjectActionProps",
		reflect.TypeOf((*S3PutObjectActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.SqsQueueAction",
		reflect.TypeOf((*SqsQueueAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_SqsQueueAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.SqsQueueActionProps",
		reflect.TypeOf((*SqsQueueActionProps)(nil)).Elem(),
	)
}

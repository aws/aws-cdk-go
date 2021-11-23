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
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.CommonActionProps",
		reflect.TypeOf((*CommonActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/aws-iot-actions-alpha.FirehoseStreamAction",
		reflect.TypeOf((*FirehoseStreamAction)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "bind", GoMethod: "Bind"},
		},
		func() interface{} {
			j := jsiiProxy_FirehoseStreamAction{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkiotalphaIAction)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/aws-iot-actions-alpha.FirehoseStreamActionProps",
		reflect.TypeOf((*FirehoseStreamActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/aws-iot-actions-alpha.FirehoseStreamRecordSeparator",
		reflect.TypeOf((*FirehoseStreamRecordSeparator)(nil)).Elem(),
		map[string]interface{}{
			"NEWLINE": FirehoseStreamRecordSeparator_NEWLINE,
			"TAB": FirehoseStreamRecordSeparator_TAB,
			"WINDOWS_NEWLINE": FirehoseStreamRecordSeparator_WINDOWS_NEWLINE,
			"COMMA": FirehoseStreamRecordSeparator_COMMA,
		},
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
}

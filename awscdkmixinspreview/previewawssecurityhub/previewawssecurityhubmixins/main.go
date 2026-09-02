package previewawssecurityhubmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubLogsMixin",
		reflect.TypeOf((*CfnHubLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHubLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogs",
		reflect.TypeOf((*CfnHubSecurityFindingLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
		},
		func() interface{} {
			return &jsiiProxy_CfnHubSecurityFindingLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsDestProps",
		reflect.TypeOf((*CfnHubSecurityFindingLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsLogGroupProps",
		reflect.TypeOf((*CfnHubSecurityFindingLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsOutputFormat",
		reflect.TypeOf((*CfnHubSecurityFindingLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnHubSecurityFindingLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnHubSecurityFindingLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnHubSecurityFindingLogsOutputFormat_LogGroup_PLAIN,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubSecurityFindingLogsRecordFields",
		reflect.TypeOf((*CfnHubSecurityFindingLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": CfnHubSecurityFindingLogsRecordFields_EVENT,
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2LogsMixin",
		reflect.TypeOf((*CfnHubV2LogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnHubV2LogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogs",
		reflect.TypeOf((*CfnHubV2SecurityFindingLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toLogGroup", GoMethod: "ToLogGroup"},
		},
		func() interface{} {
			return &jsiiProxy_CfnHubV2SecurityFindingLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsDestProps",
		reflect.TypeOf((*CfnHubV2SecurityFindingLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsLogGroupProps",
		reflect.TypeOf((*CfnHubV2SecurityFindingLogsLogGroupProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsOutputFormat",
		reflect.TypeOf((*CfnHubV2SecurityFindingLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnHubV2SecurityFindingLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsOutputFormat.LogGroup",
		reflect.TypeOf((*CfnHubV2SecurityFindingLogsOutputFormat_LogGroup)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnHubV2SecurityFindingLogsOutputFormat_LogGroup_PLAIN,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_securityhub.mixins.CfnHubV2SecurityFindingLogsRecordFields",
		reflect.TypeOf((*CfnHubV2SecurityFindingLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"EVENT": CfnHubV2SecurityFindingLogsRecordFields_EVENT,
		},
	)
}

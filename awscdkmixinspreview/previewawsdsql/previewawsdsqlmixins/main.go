package previewawsdsqlmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogs",
		reflect.TypeOf((*CfnClusterAuditLogs)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "toDestination", GoMethod: "ToDestination"},
			_jsii_.MemberMethod{JsiiMethod: "toS3", GoMethod: "ToS3"},
		},
		func() interface{} {
			return &jsiiProxy_CfnClusterAuditLogs{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsDestProps",
		reflect.TypeOf((*CfnClusterAuditLogsDestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsOutputFormat",
		reflect.TypeOf((*CfnClusterAuditLogsOutputFormat)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_CfnClusterAuditLogsOutputFormat{}
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsOutputFormat.S3",
		reflect.TypeOf((*CfnClusterAuditLogsOutputFormat_S3)(nil)).Elem(),
		map[string]interface{}{
			"PLAIN": CfnClusterAuditLogsOutputFormat_S3_PLAIN,
			"JSON": CfnClusterAuditLogsOutputFormat_S3_JSON,
			"W3C": CfnClusterAuditLogsOutputFormat_S3_W3C,
			"PARQUET": CfnClusterAuditLogsOutputFormat_S3_PARQUET,
		},
	)
	_jsii_.RegisterEnum(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsRecordFields",
		reflect.TypeOf((*CfnClusterAuditLogsRecordFields)(nil)).Elem(),
		map[string]interface{}{
			"TIMESTAMP": CfnClusterAuditLogsRecordFields_TIMESTAMP,
			"RESOURCE_ARN": CfnClusterAuditLogsRecordFields_RESOURCE_ARN,
			"EVENT_TIMESTAMP": CfnClusterAuditLogsRecordFields_EVENT_TIMESTAMP,
			"RESOURCE_ID": CfnClusterAuditLogsRecordFields_RESOURCE_ID,
			"SCHEMA_VERSION": CfnClusterAuditLogsRecordFields_SCHEMA_VERSION,
			"AWS_REGION": CfnClusterAuditLogsRecordFields_AWS_REGION,
			"CLUSTER_ID": CfnClusterAuditLogsRecordFields_CLUSTER_ID,
			"TYPE": CfnClusterAuditLogsRecordFields_TYPE,
			"DATABASE": CfnClusterAuditLogsRecordFields_DATABASE,
			"SESSION_ID": CfnClusterAuditLogsRecordFields_SESSION_ID,
			"IAM_PRINCIPAL": CfnClusterAuditLogsRecordFields_IAM_PRINCIPAL,
			"IAM_ARN": CfnClusterAuditLogsRecordFields_IAM_ARN,
			"SESSION_USER": CfnClusterAuditLogsRecordFields_SESSION_USER,
			"USER_ID": CfnClusterAuditLogsRecordFields_USER_ID,
			"CLIENT_APPLICATION": CfnClusterAuditLogsRecordFields_CLIENT_APPLICATION,
			"TRANSACTION_ID": CfnClusterAuditLogsRecordFields_TRANSACTION_ID,
			"STATEMENT_ID": CfnClusterAuditLogsRecordFields_STATEMENT_ID,
			"SUBSTATEMENT_ID": CfnClusterAuditLogsRecordFields_SUBSTATEMENT_ID,
			"CLASS": CfnClusterAuditLogsRecordFields_CLASS,
			"COMMAND": CfnClusterAuditLogsRecordFields_COMMAND,
			"OBJECT_TYPE": CfnClusterAuditLogsRecordFields_OBJECT_TYPE,
			"OBJECT_NAME": CfnClusterAuditLogsRecordFields_OBJECT_NAME,
			"STATEMENT_TEXT": CfnClusterAuditLogsRecordFields_STATEMENT_TEXT,
			"PARAMETERS": CfnClusterAuditLogsRecordFields_PARAMETERS,
			"ROWS_AFFECTED": CfnClusterAuditLogsRecordFields_ROWS_AFFECTED,
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterAuditLogsS3Props",
		reflect.TypeOf((*CfnClusterAuditLogsS3Props)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_dsql.mixins.CfnClusterLogsMixin",
		reflect.TypeOf((*CfnClusterLogsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "logDelivery", GoGetter: "LogDelivery"},
			_jsii_.MemberProperty{JsiiProperty: "logType", GoGetter: "LogType"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnClusterLogsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}

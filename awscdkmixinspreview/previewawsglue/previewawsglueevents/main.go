package previewawsglueevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents",
		reflect.TypeOf((*DatabaseEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "glueDataCatalogDatabaseStateChangePattern", GoMethod: "GlueDataCatalogDatabaseStateChangePattern"},
			_jsii_.MemberMethod{JsiiMethod: "glueDataCatalogTableStateChangePattern", GoMethod: "GlueDataCatalogTableStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_DatabaseEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogDatabaseStateChange",
		reflect.TypeOf((*GlueDataCatalogDatabaseStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GlueDataCatalogDatabaseStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogDatabaseStateChange.GlueDataCatalogDatabaseStateChangeProps",
		reflect.TypeOf((*GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogTableStateChange",
		reflect.TypeOf((*GlueDataCatalogTableStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GlueDataCatalogTableStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueDataCatalogTableStateChange.GlueDataCatalogTableStateChangeProps",
		reflect.TypeOf((*GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobRunStatus",
		reflect.TypeOf((*GlueJobRunStatus)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GlueJobRunStatus{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobRunStatus.GlueJobRunStatusProps",
		reflect.TypeOf((*GlueJobRunStatus_GlueJobRunStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobRunStatus.NotificationCondition",
		reflect.TypeOf((*GlueJobRunStatus_NotificationCondition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobStateChange",
		reflect.TypeOf((*GlueJobStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_GlueJobStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.GlueJobStateChange.GlueJobStateChangeProps",
		reflect.TypeOf((*GlueJobStateChange_GlueJobStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents",
		reflect.TypeOf((*JobEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "glueJobRunStatusPattern", GoMethod: "GlueJobRunStatusPattern"},
			_jsii_.MemberMethod{JsiiMethod: "glueJobStateChangePattern", GoMethod: "GlueJobStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_JobEvents{}
		},
	)
}

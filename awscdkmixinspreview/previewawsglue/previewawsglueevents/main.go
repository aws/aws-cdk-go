package previewawsglueevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
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
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogDatabaseStateChange",
		reflect.TypeOf((*DatabaseEvents_GlueDataCatalogDatabaseStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatabaseEvents_GlueDataCatalogDatabaseStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogDatabaseStateChange.GlueDataCatalogDatabaseStateChangeProps",
		reflect.TypeOf((*DatabaseEvents_GlueDataCatalogDatabaseStateChange_GlueDataCatalogDatabaseStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogTableStateChange",
		reflect.TypeOf((*DatabaseEvents_GlueDataCatalogTableStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_DatabaseEvents_GlueDataCatalogTableStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.DatabaseEvents.GlueDataCatalogTableStateChange.GlueDataCatalogTableStateChangeProps",
		reflect.TypeOf((*DatabaseEvents_GlueDataCatalogTableStateChange_GlueDataCatalogTableStateChangeProps)(nil)).Elem(),
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
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_JobEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.ResponseElements",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_ResponseElements)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*JobEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobRunStatus",
		reflect.TypeOf((*JobEvents_GlueJobRunStatus)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_JobEvents_GlueJobRunStatus{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobRunStatus.GlueJobRunStatusProps",
		reflect.TypeOf((*JobEvents_GlueJobRunStatus_GlueJobRunStatusProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobRunStatus.NotificationCondition",
		reflect.TypeOf((*JobEvents_GlueJobRunStatus_NotificationCondition)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobStateChange",
		reflect.TypeOf((*JobEvents_GlueJobStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_JobEvents_GlueJobStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_glue.events.JobEvents.GlueJobStateChange.GlueJobStateChangeProps",
		reflect.TypeOf((*JobEvents_GlueJobStateChange_GlueJobStateChangeProps)(nil)).Elem(),
	)
}

package previewawslogsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents",
		reflect.TypeOf((*LogGroupEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
		},
		func() interface{} {
			return &jsiiProxy_LogGroupEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_LogGroupEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_logs.events.LogGroupEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*LogGroupEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
}

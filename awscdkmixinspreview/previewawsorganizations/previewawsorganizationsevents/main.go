package previewawsorganizationsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents",
		reflect.TypeOf((*AccountEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsServiceEventViaCloudTrailPattern", GoMethod: "AwsServiceEventViaCloudTrailPattern"},
		},
		func() interface{} {
			return &jsiiProxy_AccountEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail",
		reflect.TypeOf((*AccountEvents_AWSServiceEventViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AccountEvents_AWSServiceEventViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail.AWSServiceEventViaCloudTrailProps",
		reflect.TypeOf((*AccountEvents_AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail.CreateAccountStatus",
		reflect.TypeOf((*AccountEvents_AWSServiceEventViaCloudTrail_CreateAccountStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail.ServiceEventDetails",
		reflect.TypeOf((*AccountEvents_AWSServiceEventViaCloudTrail_ServiceEventDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AccountEvents.AWSServiceEventViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AccountEvents_AWSServiceEventViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
}

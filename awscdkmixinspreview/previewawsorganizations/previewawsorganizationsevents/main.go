package previewawsorganizationsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSServiceEventViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail.AWSServiceEventViaCloudTrailProps",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_AWSServiceEventViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail.CreateAccountStatus",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_CreateAccountStatus)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail.ServiceEventDetails",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_ServiceEventDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_organizations.events.AWSServiceEventViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSServiceEventViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
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
}

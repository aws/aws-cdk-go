package previewawscodecommitevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents",
		reflect.TypeOf((*RepositoryEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "codeCommitCommentOnCommitPattern", GoMethod: "CodeCommitCommentOnCommitPattern"},
			_jsii_.MemberMethod{JsiiMethod: "codeCommitCommentOnPullRequestPattern", GoMethod: "CodeCommitCommentOnPullRequestPattern"},
			_jsii_.MemberMethod{JsiiMethod: "codeCommitRepositoryStateChangePattern", GoMethod: "CodeCommitRepositoryStateChangePattern"},
		},
		func() interface{} {
			return &jsiiProxy_RepositoryEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.AdditionalEventData",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_AdditionalEventData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.Location",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_Location)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.RequestParametersItem1",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem1)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.AWSAPICallViaCloudTrail.RequestParametersItem2",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem2)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnCommit",
		reflect.TypeOf((*RepositoryEvents_CodeCommitCommentOnCommit)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_CodeCommitCommentOnCommit{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnCommit.CodeCommitCommentOnCommitProps",
		reflect.TypeOf((*RepositoryEvents_CodeCommitCommentOnCommit_CodeCommitCommentOnCommitProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnPullRequest",
		reflect.TypeOf((*RepositoryEvents_CodeCommitCommentOnPullRequest)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_CodeCommitCommentOnPullRequest{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitCommentOnPullRequest.CodeCommitCommentOnPullRequestProps",
		reflect.TypeOf((*RepositoryEvents_CodeCommitCommentOnPullRequest_CodeCommitCommentOnPullRequestProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitRepositoryStateChange",
		reflect.TypeOf((*RepositoryEvents_CodeCommitRepositoryStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_CodeCommitRepositoryStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_codecommit.events.RepositoryEvents.CodeCommitRepositoryStateChange.CodeCommitRepositoryStateChangeProps",
		reflect.TypeOf((*RepositoryEvents_CodeCommitRepositoryStateChange_CodeCommitRepositoryStateChangeProps)(nil)).Elem(),
	)
}

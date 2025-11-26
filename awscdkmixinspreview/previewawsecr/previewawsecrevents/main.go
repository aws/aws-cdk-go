package previewawsecrevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents",
		reflect.TypeOf((*RepositoryEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCRImageActionPattern", GoMethod: "ECRImageActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCRImageScanPattern", GoMethod: "ECRImageScanPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCRPullThroughCacheActionPattern", GoMethod: "ECRPullThroughCacheActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCRReferrerActionPattern", GoMethod: "ECRReferrerActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "eCRReplicationActionPattern", GoMethod: "ECRReplicationActionPattern"},
		},
		func() interface{} {
			return &jsiiProxy_RepositoryEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.AwsapiCallViaCloudTrailItem",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*RepositoryEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageAction",
		reflect.TypeOf((*RepositoryEvents_ECRImageAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_ECRImageAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageAction.ECRImageActionProps",
		reflect.TypeOf((*RepositoryEvents_ECRImageAction_ECRImageActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageScan",
		reflect.TypeOf((*RepositoryEvents_ECRImageScan)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_ECRImageScan{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageScan.ECRImageScanProps",
		reflect.TypeOf((*RepositoryEvents_ECRImageScan_ECRImageScanProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRImageScan.FindingSeverityCounts",
		reflect.TypeOf((*RepositoryEvents_ECRImageScan_FindingSeverityCounts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRPullThroughCacheAction",
		reflect.TypeOf((*RepositoryEvents_ECRPullThroughCacheAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_ECRPullThroughCacheAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRPullThroughCacheAction.ECRPullThroughCacheActionProps",
		reflect.TypeOf((*RepositoryEvents_ECRPullThroughCacheAction_ECRPullThroughCacheActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReferrerAction",
		reflect.TypeOf((*RepositoryEvents_ECRReferrerAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_ECRReferrerAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReferrerAction.ECRReferrerActionProps",
		reflect.TypeOf((*RepositoryEvents_ECRReferrerAction_ECRReferrerActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReplicationAction",
		reflect.TypeOf((*RepositoryEvents_ECRReplicationAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_RepositoryEvents_ECRReplicationAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents.ECRReplicationAction.ECRReplicationActionProps",
		reflect.TypeOf((*RepositoryEvents_ECRReplicationAction_ECRReplicationActionProps)(nil)).Elem(),
	)
}

package previewawsecrevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.AwsapiCallViaCloudTrailItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.RequestParametersItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParametersItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageAction",
		reflect.TypeOf((*ECRImageAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECRImageAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageAction.ECRImageActionProps",
		reflect.TypeOf((*ECRImageAction_ECRImageActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageScan",
		reflect.TypeOf((*ECRImageScan)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECRImageScan{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageScan.ECRImageScanProps",
		reflect.TypeOf((*ECRImageScan_ECRImageScanProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRImageScan.FindingSeverityCounts",
		reflect.TypeOf((*ECRImageScan_FindingSeverityCounts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRPullThroughCacheAction",
		reflect.TypeOf((*ECRPullThroughCacheAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECRPullThroughCacheAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRPullThroughCacheAction.ECRPullThroughCacheActionProps",
		reflect.TypeOf((*ECRPullThroughCacheAction_ECRPullThroughCacheActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReferrerAction",
		reflect.TypeOf((*ECRReferrerAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECRReferrerAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReferrerAction.ECRReferrerActionProps",
		reflect.TypeOf((*ECRReferrerAction_ECRReferrerActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReplicationAction",
		reflect.TypeOf((*ECRReplicationAction)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ECRReplicationAction{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ecr.events.ECRReplicationAction.ECRReplicationActionProps",
		reflect.TypeOf((*ECRReplicationAction_ECRReplicationActionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ecr.events.RepositoryEvents",
		reflect.TypeOf((*RepositoryEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecrImageActionPattern", GoMethod: "EcrImageActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecrImageScanPattern", GoMethod: "EcrImageScanPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecrPullThroughCacheActionPattern", GoMethod: "EcrPullThroughCacheActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecrReferrerActionPattern", GoMethod: "EcrReferrerActionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "ecrReplicationActionPattern", GoMethod: "EcrReplicationActionPattern"},
		},
		func() interface{} {
			return &jsiiProxy_RepositoryEvents{}
		},
	)
}

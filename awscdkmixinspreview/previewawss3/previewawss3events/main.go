package previewawss3events

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.AdditionalEventData",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AdditionalEventData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.AwsapiCallViaCloudTrailItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.LegalHoldInfo",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_LegalHoldInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.ObjectRetentionInfo",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_ObjectRetentionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.RetentionInfo",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RetentionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.TlsDetails",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_TlsDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents",
		reflect.TypeOf((*BucketEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectAccessTierChangedPattern", GoMethod: "ObjectAccessTierChangedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectACLUpdatedPattern", GoMethod: "ObjectACLUpdatedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectCreatedPattern", GoMethod: "ObjectCreatedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectDeletedPattern", GoMethod: "ObjectDeletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectRestoreCompletedPattern", GoMethod: "ObjectRestoreCompletedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectRestoreExpiredPattern", GoMethod: "ObjectRestoreExpiredPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectRestoreInitiatedPattern", GoMethod: "ObjectRestoreInitiatedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectStorageClassChangedPattern", GoMethod: "ObjectStorageClassChangedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectTagsAddedPattern", GoMethod: "ObjectTagsAddedPattern"},
			_jsii_.MemberMethod{JsiiMethod: "objectTagsDeletedPattern", GoMethod: "ObjectTagsDeletedPattern"},
		},
		func() interface{} {
			return &jsiiProxy_BucketEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated",
		reflect.TypeOf((*ObjectACLUpdated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectACLUpdated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated.Bucket",
		reflect.TypeOf((*ObjectACLUpdated_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated.ObjectACLUpdatedProps",
		reflect.TypeOf((*ObjectACLUpdated_ObjectACLUpdatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectACLUpdated.ObjectType",
		reflect.TypeOf((*ObjectACLUpdated_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged",
		reflect.TypeOf((*ObjectAccessTierChanged)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectAccessTierChanged{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged.Bucket",
		reflect.TypeOf((*ObjectAccessTierChanged_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged.ObjectAccessTierChangedProps",
		reflect.TypeOf((*ObjectAccessTierChanged_ObjectAccessTierChangedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectAccessTierChanged.ObjectType",
		reflect.TypeOf((*ObjectAccessTierChanged_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated",
		reflect.TypeOf((*ObjectCreated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectCreated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated.Bucket",
		reflect.TypeOf((*ObjectCreated_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated.ObjectCreatedProps",
		reflect.TypeOf((*ObjectCreated_ObjectCreatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectCreated.ObjectType",
		reflect.TypeOf((*ObjectCreated_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted",
		reflect.TypeOf((*ObjectDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted.Bucket",
		reflect.TypeOf((*ObjectDeleted_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted.ObjectDeletedProps",
		reflect.TypeOf((*ObjectDeleted_ObjectDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectDeleted.ObjectType",
		reflect.TypeOf((*ObjectDeleted_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted",
		reflect.TypeOf((*ObjectRestoreCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectRestoreCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted.Bucket",
		reflect.TypeOf((*ObjectRestoreCompleted_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted.ObjectRestoreCompletedProps",
		reflect.TypeOf((*ObjectRestoreCompleted_ObjectRestoreCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreCompleted.ObjectType",
		reflect.TypeOf((*ObjectRestoreCompleted_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired",
		reflect.TypeOf((*ObjectRestoreExpired)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectRestoreExpired{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired.Bucket",
		reflect.TypeOf((*ObjectRestoreExpired_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired.ObjectRestoreExpiredProps",
		reflect.TypeOf((*ObjectRestoreExpired_ObjectRestoreExpiredProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreExpired.ObjectType",
		reflect.TypeOf((*ObjectRestoreExpired_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated",
		reflect.TypeOf((*ObjectRestoreInitiated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectRestoreInitiated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated.Bucket",
		reflect.TypeOf((*ObjectRestoreInitiated_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated.ObjectRestoreInitiatedProps",
		reflect.TypeOf((*ObjectRestoreInitiated_ObjectRestoreInitiatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectRestoreInitiated.ObjectType",
		reflect.TypeOf((*ObjectRestoreInitiated_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged",
		reflect.TypeOf((*ObjectStorageClassChanged)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectStorageClassChanged{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged.Bucket",
		reflect.TypeOf((*ObjectStorageClassChanged_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged.ObjectStorageClassChangedProps",
		reflect.TypeOf((*ObjectStorageClassChanged_ObjectStorageClassChangedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectStorageClassChanged.ObjectType",
		reflect.TypeOf((*ObjectStorageClassChanged_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded",
		reflect.TypeOf((*ObjectTagsAdded)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectTagsAdded{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded.Bucket",
		reflect.TypeOf((*ObjectTagsAdded_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded.ObjectTagsAddedProps",
		reflect.TypeOf((*ObjectTagsAdded_ObjectTagsAddedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsAdded.ObjectType",
		reflect.TypeOf((*ObjectTagsAdded_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted",
		reflect.TypeOf((*ObjectTagsDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_ObjectTagsDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted.Bucket",
		reflect.TypeOf((*ObjectTagsDeleted_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted.ObjectTagsDeletedProps",
		reflect.TypeOf((*ObjectTagsDeleted_ObjectTagsDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.ObjectTagsDeleted.ObjectType",
		reflect.TypeOf((*ObjectTagsDeleted_ObjectType)(nil)).Elem(),
	)
}

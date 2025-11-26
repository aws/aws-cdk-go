package previewawss3events

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
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
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.AdditionalEventData",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_AdditionalEventData)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.AwsapiCallViaCloudTrailItem",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.LegalHoldInfo",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_LegalHoldInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.ObjectRetentionInfo",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_ObjectRetentionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.RetentionInfo",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_RetentionInfo)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.TlsDetails",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_TlsDetails)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*BucketEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectACLUpdated",
		reflect.TypeOf((*BucketEvents_ObjectACLUpdated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectACLUpdated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectACLUpdated.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectACLUpdated_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectACLUpdated.ObjectACLUpdatedProps",
		reflect.TypeOf((*BucketEvents_ObjectACLUpdated_ObjectACLUpdatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectACLUpdated.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectACLUpdated_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectAccessTierChanged",
		reflect.TypeOf((*BucketEvents_ObjectAccessTierChanged)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectAccessTierChanged{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectAccessTierChanged.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectAccessTierChanged_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectAccessTierChanged.ObjectAccessTierChangedProps",
		reflect.TypeOf((*BucketEvents_ObjectAccessTierChanged_ObjectAccessTierChangedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectAccessTierChanged.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectAccessTierChanged_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectCreated",
		reflect.TypeOf((*BucketEvents_ObjectCreated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectCreated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectCreated.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectCreated_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectCreated.ObjectCreatedProps",
		reflect.TypeOf((*BucketEvents_ObjectCreated_ObjectCreatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectCreated.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectCreated_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectDeleted",
		reflect.TypeOf((*BucketEvents_ObjectDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectDeleted.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectDeleted_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectDeleted.ObjectDeletedProps",
		reflect.TypeOf((*BucketEvents_ObjectDeleted_ObjectDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectDeleted.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectDeleted_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreCompleted",
		reflect.TypeOf((*BucketEvents_ObjectRestoreCompleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectRestoreCompleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreCompleted.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectRestoreCompleted_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreCompleted.ObjectRestoreCompletedProps",
		reflect.TypeOf((*BucketEvents_ObjectRestoreCompleted_ObjectRestoreCompletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreCompleted.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectRestoreCompleted_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreExpired",
		reflect.TypeOf((*BucketEvents_ObjectRestoreExpired)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectRestoreExpired{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreExpired.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectRestoreExpired_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreExpired.ObjectRestoreExpiredProps",
		reflect.TypeOf((*BucketEvents_ObjectRestoreExpired_ObjectRestoreExpiredProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreExpired.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectRestoreExpired_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreInitiated",
		reflect.TypeOf((*BucketEvents_ObjectRestoreInitiated)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectRestoreInitiated{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreInitiated.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectRestoreInitiated_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreInitiated.ObjectRestoreInitiatedProps",
		reflect.TypeOf((*BucketEvents_ObjectRestoreInitiated_ObjectRestoreInitiatedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectRestoreInitiated.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectRestoreInitiated_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectStorageClassChanged",
		reflect.TypeOf((*BucketEvents_ObjectStorageClassChanged)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectStorageClassChanged{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectStorageClassChanged.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectStorageClassChanged_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectStorageClassChanged.ObjectStorageClassChangedProps",
		reflect.TypeOf((*BucketEvents_ObjectStorageClassChanged_ObjectStorageClassChangedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectStorageClassChanged.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectStorageClassChanged_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsAdded",
		reflect.TypeOf((*BucketEvents_ObjectTagsAdded)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectTagsAdded{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsAdded.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectTagsAdded_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsAdded.ObjectTagsAddedProps",
		reflect.TypeOf((*BucketEvents_ObjectTagsAdded_ObjectTagsAddedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsAdded.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectTagsAdded_ObjectType)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsDeleted",
		reflect.TypeOf((*BucketEvents_ObjectTagsDeleted)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_BucketEvents_ObjectTagsDeleted{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsDeleted.Bucket",
		reflect.TypeOf((*BucketEvents_ObjectTagsDeleted_Bucket)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsDeleted.ObjectTagsDeletedProps",
		reflect.TypeOf((*BucketEvents_ObjectTagsDeleted_ObjectTagsDeletedProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_s3.events.BucketEvents.ObjectTagsDeleted.ObjectType",
		reflect.TypeOf((*BucketEvents_ObjectTagsDeleted_ObjectType)(nil)).Elem(),
	)
}

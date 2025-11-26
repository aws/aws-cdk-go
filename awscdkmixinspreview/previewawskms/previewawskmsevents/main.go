package previewawskmsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents",
		reflect.TypeOf((*KeyEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "kMSCMKDeletionPattern", GoMethod: "KMSCMKDeletionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "kMSCMKRotationPattern", GoMethod: "KMSCMKRotationPattern"},
			_jsii_.MemberMethod{JsiiMethod: "kMSImportedKeyMaterialExpirationPattern", GoMethod: "KMSImportedKeyMaterialExpirationPattern"},
		},
		func() interface{} {
			return &jsiiProxy_KeyEvents{}
		},
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KeyEvents_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.AwsapiCallViaCloudTrailItem",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.EncryptionContext",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_EncryptionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*KeyEvents_AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKDeletion",
		reflect.TypeOf((*KeyEvents_KMSCMKDeletion)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KeyEvents_KMSCMKDeletion{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKDeletion.KMSCMKDeletionProps",
		reflect.TypeOf((*KeyEvents_KMSCMKDeletion_KMSCMKDeletionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKRotation",
		reflect.TypeOf((*KeyEvents_KMSCMKRotation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KeyEvents_KMSCMKRotation{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSCMKRotation.KMSCMKRotationProps",
		reflect.TypeOf((*KeyEvents_KMSCMKRotation_KMSCMKRotationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSImportedKeyMaterialExpiration",
		reflect.TypeOf((*KeyEvents_KMSImportedKeyMaterialExpiration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KeyEvents_KMSImportedKeyMaterialExpiration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents.KMSImportedKeyMaterialExpiration.KMSImportedKeyMaterialExpirationProps",
		reflect.TypeOf((*KeyEvents_KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps)(nil)).Elem(),
	)
}

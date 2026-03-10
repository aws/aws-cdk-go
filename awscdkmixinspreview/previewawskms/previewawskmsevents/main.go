package previewawskmsevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail",
		reflect.TypeOf((*AWSAPICallViaCloudTrail)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_AWSAPICallViaCloudTrail{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.AWSAPICallViaCloudTrailProps",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AWSAPICallViaCloudTrailProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.Attributes",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_Attributes)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.AwsapiCallViaCloudTrailItem",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_AwsapiCallViaCloudTrailItem)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.EncryptionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_EncryptionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.RequestParameters",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_RequestParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.SessionContext",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionContext)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.SessionIssuer",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_SessionIssuer)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.AWSAPICallViaCloudTrail.UserIdentity",
		reflect.TypeOf((*AWSAPICallViaCloudTrail_UserIdentity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKDeletion",
		reflect.TypeOf((*KMSCMKDeletion)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KMSCMKDeletion{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKDeletion.KMSCMKDeletionProps",
		reflect.TypeOf((*KMSCMKDeletion_KMSCMKDeletionProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKRotation",
		reflect.TypeOf((*KMSCMKRotation)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KMSCMKRotation{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSCMKRotation.KMSCMKRotationProps",
		reflect.TypeOf((*KMSCMKRotation_KMSCMKRotationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSImportedKeyMaterialExpiration",
		reflect.TypeOf((*KMSImportedKeyMaterialExpiration)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_KMSImportedKeyMaterialExpiration{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kms.events.KMSImportedKeyMaterialExpiration.KMSImportedKeyMaterialExpirationProps",
		reflect.TypeOf((*KMSImportedKeyMaterialExpiration_KMSImportedKeyMaterialExpirationProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kms.events.KeyEvents",
		reflect.TypeOf((*KeyEvents)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "awsAPICallViaCloudTrailPattern", GoMethod: "AwsAPICallViaCloudTrailPattern"},
			_jsii_.MemberMethod{JsiiMethod: "kmsCMKDeletionPattern", GoMethod: "KmsCMKDeletionPattern"},
			_jsii_.MemberMethod{JsiiMethod: "kmsCMKRotationPattern", GoMethod: "KmsCMKRotationPattern"},
			_jsii_.MemberMethod{JsiiMethod: "kmsImportedKeyMaterialExpirationPattern", GoMethod: "KmsImportedKeyMaterialExpirationPattern"},
		},
		func() interface{} {
			return &jsiiProxy_KeyEvents{}
		},
	)
}

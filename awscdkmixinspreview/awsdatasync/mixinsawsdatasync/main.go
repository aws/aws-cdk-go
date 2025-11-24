package mixinsawsdatasync

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnAgentMixinProps",
		reflect.TypeOf((*CfnAgentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnAgentPropsMixin",
		reflect.TypeOf((*CfnAgentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationAzureBlobMixinProps",
		reflect.TypeOf((*CfnLocationAzureBlobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationAzureBlobPropsMixin",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationAzureBlobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationAzureBlobPropsMixin.AzureBlobSasConfigurationProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_AzureBlobSasConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationAzureBlobPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationAzureBlobPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationAzureBlobPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSMixinProps",
		reflect.TypeOf((*CfnLocationEFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSPropsMixin",
		reflect.TypeOf((*CfnLocationEFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationEFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationEFSPropsMixin.Ec2ConfigProperty",
		reflect.TypeOf((*CfnLocationEFSPropsMixin_Ec2ConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxLustreMixinProps",
		reflect.TypeOf((*CfnLocationFSxLustreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxLustrePropsMixin",
		reflect.TypeOf((*CfnLocationFSxLustrePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxLustrePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPMixinProps",
		reflect.TypeOf((*CfnLocationFSxONTAPMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxONTAPPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin.NFSProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_NFSProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin.NfsMountOptionsProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_NfsMountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin.ProtocolProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_ProtocolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin.SMBProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_SMBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxONTAPPropsMixin.SmbMountOptionsProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_SmbMountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxOpenZFSMixinProps",
		reflect.TypeOf((*CfnLocationFSxOpenZFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxOpenZFSPropsMixin",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxOpenZFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxOpenZFSPropsMixin.MountOptionsProperty",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin_MountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxOpenZFSPropsMixin.NFSProperty",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin_NFSProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxOpenZFSPropsMixin.ProtocolProperty",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin_ProtocolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxWindowsMixinProps",
		reflect.TypeOf((*CfnLocationFSxWindowsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationFSxWindowsPropsMixin",
		reflect.TypeOf((*CfnLocationFSxWindowsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxWindowsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSMixinProps",
		reflect.TypeOf((*CfnLocationHDFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationHDFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin.NameNodeProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_NameNodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationHDFSPropsMixin.QopConfigurationProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_QopConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationNFSMixinProps",
		reflect.TypeOf((*CfnLocationNFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationNFSPropsMixin",
		reflect.TypeOf((*CfnLocationNFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationNFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationNFSPropsMixin.MountOptionsProperty",
		reflect.TypeOf((*CfnLocationNFSPropsMixin_MountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationNFSPropsMixin.OnPremConfigProperty",
		reflect.TypeOf((*CfnLocationNFSPropsMixin_OnPremConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStorageMixinProps",
		reflect.TypeOf((*CfnLocationObjectStorageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationObjectStoragePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationObjectStoragePropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationS3MixinProps",
		reflect.TypeOf((*CfnLocationS3MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationS3PropsMixin",
		reflect.TypeOf((*CfnLocationS3PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationS3PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationS3PropsMixin.S3ConfigProperty",
		reflect.TypeOf((*CfnLocationS3PropsMixin_S3ConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationSMBMixinProps",
		reflect.TypeOf((*CfnLocationSMBMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationSMBPropsMixin",
		reflect.TypeOf((*CfnLocationSMBPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationSMBPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationSMBPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationSMBPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationSMBPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnLocationSMBPropsMixin.MountOptionsProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_MountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskMixinProps",
		reflect.TypeOf((*CfnTaskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin",
		reflect.TypeOf((*CfnTaskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTaskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.DeletedProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_DeletedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.FilterRuleProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_FilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.ManifestConfigProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_ManifestConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.ManifestConfigSourceS3Property",
		reflect.TypeOf((*CfnTaskPropsMixin_ManifestConfigSourceS3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.OptionsProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_OptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.OverridesProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_OverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.S3Property",
		reflect.TypeOf((*CfnTaskPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.SkippedProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_SkippedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.TaskReportConfigProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_TaskReportConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.TaskScheduleProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_TaskScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.TransferredProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_TransferredProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_datasync.mixins.CfnTaskPropsMixin.VerifiedProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_VerifiedProperty)(nil)).Elem(),
	)
}

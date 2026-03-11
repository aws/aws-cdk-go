package awsdatasync

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnAgentMixinProps",
		reflect.TypeOf((*CfnAgentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnAgentPropsMixin",
		reflect.TypeOf((*CfnAgentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAgentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobMixinProps",
		reflect.TypeOf((*CfnLocationAzureBlobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationAzureBlobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin.AzureBlobSasConfigurationProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_AzureBlobSasConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationAzureBlobPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationAzureBlobPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationEFSMixinProps",
		reflect.TypeOf((*CfnLocationEFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationEFSPropsMixin",
		reflect.TypeOf((*CfnLocationEFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationEFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationEFSPropsMixin.Ec2ConfigProperty",
		reflect.TypeOf((*CfnLocationEFSPropsMixin_Ec2ConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxLustreMixinProps",
		reflect.TypeOf((*CfnLocationFSxLustreMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxLustrePropsMixin",
		reflect.TypeOf((*CfnLocationFSxLustrePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxLustrePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPMixinProps",
		reflect.TypeOf((*CfnLocationFSxONTAPMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxONTAPPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.NFSProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_NFSProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.NfsMountOptionsProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_NfsMountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.ProtocolProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_ProtocolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.SMBProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_SMBProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxONTAPPropsMixin.SmbMountOptionsProperty",
		reflect.TypeOf((*CfnLocationFSxONTAPPropsMixin_SmbMountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxOpenZFSMixinProps",
		reflect.TypeOf((*CfnLocationFSxOpenZFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxOpenZFSPropsMixin",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxOpenZFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxOpenZFSPropsMixin.MountOptionsProperty",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin_MountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxOpenZFSPropsMixin.NFSProperty",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin_NFSProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxOpenZFSPropsMixin.ProtocolProperty",
		reflect.TypeOf((*CfnLocationFSxOpenZFSPropsMixin_ProtocolProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsMixinProps",
		reflect.TypeOf((*CfnLocationFSxWindowsMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin",
		reflect.TypeOf((*CfnLocationFSxWindowsPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationFSxWindowsPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationFSxWindowsPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationFSxWindowsPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationFSxWindowsPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationFSxWindowsPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSMixinProps",
		reflect.TypeOf((*CfnLocationHDFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSPropsMixin",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationHDFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSPropsMixin.NameNodeProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_NameNodeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationHDFSPropsMixin.QopConfigurationProperty",
		reflect.TypeOf((*CfnLocationHDFSPropsMixin_QopConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSMixinProps",
		reflect.TypeOf((*CfnLocationNFSMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin",
		reflect.TypeOf((*CfnLocationNFSPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationNFSPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin.MountOptionsProperty",
		reflect.TypeOf((*CfnLocationNFSPropsMixin_MountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationNFSPropsMixin.OnPremConfigProperty",
		reflect.TypeOf((*CfnLocationNFSPropsMixin_OnPremConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationObjectStorageMixinProps",
		reflect.TypeOf((*CfnLocationObjectStorageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationObjectStoragePropsMixin",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationObjectStoragePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationObjectStoragePropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationObjectStoragePropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationObjectStoragePropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationObjectStoragePropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationS3MixinProps",
		reflect.TypeOf((*CfnLocationS3MixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationS3PropsMixin",
		reflect.TypeOf((*CfnLocationS3PropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationS3PropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationS3PropsMixin.S3ConfigProperty",
		reflect.TypeOf((*CfnLocationS3PropsMixin_S3ConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationSMBMixinProps",
		reflect.TypeOf((*CfnLocationSMBMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationSMBPropsMixin",
		reflect.TypeOf((*CfnLocationSMBPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLocationSMBPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationSMBPropsMixin.CmkSecretConfigProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_CmkSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationSMBPropsMixin.CustomSecretConfigProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_CustomSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationSMBPropsMixin.ManagedSecretConfigProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_ManagedSecretConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnLocationSMBPropsMixin.MountOptionsProperty",
		reflect.TypeOf((*CfnLocationSMBPropsMixin_MountOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskMixinProps",
		reflect.TypeOf((*CfnTaskMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin",
		reflect.TypeOf((*CfnTaskPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTaskPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.DeletedProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_DeletedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.DestinationProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_DestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.FilterRuleProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_FilterRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.ManifestConfigProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_ManifestConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.ManifestConfigSourceS3Property",
		reflect.TypeOf((*CfnTaskPropsMixin_ManifestConfigSourceS3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.OptionsProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_OptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.OverridesProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_OverridesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.S3Property",
		reflect.TypeOf((*CfnTaskPropsMixin_S3Property)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.SkippedProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_SkippedProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.SourceProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_SourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.TaskReportConfigProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_TaskReportConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.TaskScheduleProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_TaskScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.TransferredProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_TransferredProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_datasync.CfnTaskPropsMixin.VerifiedProperty",
		reflect.TypeOf((*CfnTaskPropsMixin_VerifiedProperty)(nil)).Elem(),
	)
}

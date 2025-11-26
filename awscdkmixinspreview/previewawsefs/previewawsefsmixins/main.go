package previewawsefsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointMixinProps",
		reflect.TypeOf((*CfnAccessPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin",
		reflect.TypeOf((*CfnAccessPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin.AccessPointTagProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_AccessPointTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin.CreationInfoProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_CreationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin.PosixUserProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_PosixUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnAccessPointPropsMixin.RootDirectoryProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_RootDirectoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemMixinProps",
		reflect.TypeOf((*CfnFileSystemMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin",
		reflect.TypeOf((*CfnFileSystemPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFileSystemPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin.BackupPolicyProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_BackupPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin.ElasticFileSystemTagProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ElasticFileSystemTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin.FileSystemProtectionProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_FileSystemProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin.LifecyclePolicyProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_LifecyclePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin.ReplicationConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ReplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnFileSystemPropsMixin.ReplicationDestinationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ReplicationDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnMountTargetMixinProps",
		reflect.TypeOf((*CfnMountTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_efs.mixins.CfnMountTargetPropsMixin",
		reflect.TypeOf((*CfnMountTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMountTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}

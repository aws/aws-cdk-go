package awsefs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnAccessPointMixinProps",
		reflect.TypeOf((*CfnAccessPointMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnAccessPointPropsMixin",
		reflect.TypeOf((*CfnAccessPointPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAccessPointPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnAccessPointPropsMixin.AccessPointTagProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_AccessPointTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnAccessPointPropsMixin.CreationInfoProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_CreationInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnAccessPointPropsMixin.PosixUserProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_PosixUserProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnAccessPointPropsMixin.RootDirectoryProperty",
		reflect.TypeOf((*CfnAccessPointPropsMixin_RootDirectoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemMixinProps",
		reflect.TypeOf((*CfnFileSystemMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin",
		reflect.TypeOf((*CfnFileSystemPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnFileSystemPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin.BackupPolicyProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_BackupPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin.ElasticFileSystemTagProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ElasticFileSystemTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin.FileSystemProtectionProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_FileSystemProtectionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin.LifecyclePolicyProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_LifecyclePolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin.ReplicationConfigurationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ReplicationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnFileSystemPropsMixin.ReplicationDestinationProperty",
		reflect.TypeOf((*CfnFileSystemPropsMixin_ReplicationDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnMountTargetMixinProps",
		reflect.TypeOf((*CfnMountTargetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_efs.CfnMountTargetPropsMixin",
		reflect.TypeOf((*CfnMountTargetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMountTargetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}

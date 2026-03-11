package awsnimblestudio

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfileMixinProps",
		reflect.TypeOf((*CfnLaunchProfileMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfilePropsMixin",
		reflect.TypeOf((*CfnLaunchProfilePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnLaunchProfilePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfilePropsMixin.StreamConfigurationProperty",
		reflect.TypeOf((*CfnLaunchProfilePropsMixin_StreamConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfilePropsMixin.StreamConfigurationSessionBackupProperty",
		reflect.TypeOf((*CfnLaunchProfilePropsMixin_StreamConfigurationSessionBackupProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfilePropsMixin.StreamConfigurationSessionStorageProperty",
		reflect.TypeOf((*CfnLaunchProfilePropsMixin_StreamConfigurationSessionStorageProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfilePropsMixin.StreamingSessionStorageRootProperty",
		reflect.TypeOf((*CfnLaunchProfilePropsMixin_StreamingSessionStorageRootProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnLaunchProfilePropsMixin.VolumeConfigurationProperty",
		reflect.TypeOf((*CfnLaunchProfilePropsMixin_VolumeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStreamingImageMixinProps",
		reflect.TypeOf((*CfnStreamingImageMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStreamingImagePropsMixin",
		reflect.TypeOf((*CfnStreamingImagePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStreamingImagePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStreamingImagePropsMixin.StreamingImageEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnStreamingImagePropsMixin_StreamingImageEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentMixinProps",
		reflect.TypeOf((*CfnStudioComponentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin",
		reflect.TypeOf((*CfnStudioComponentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioComponentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.ActiveDirectoryComputerAttributeProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_ActiveDirectoryComputerAttributeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.ActiveDirectoryConfigurationProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_ActiveDirectoryConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.ComputeFarmConfigurationProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_ComputeFarmConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.LicenseServiceConfigurationProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_LicenseServiceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.ScriptParameterKeyValueProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_ScriptParameterKeyValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.SharedFileSystemConfigurationProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_SharedFileSystemConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.StudioComponentConfigurationProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_StudioComponentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioComponentPropsMixin.StudioComponentInitializationScriptProperty",
		reflect.TypeOf((*CfnStudioComponentPropsMixin_StudioComponentInitializationScriptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioMixinProps",
		reflect.TypeOf((*CfnStudioMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioPropsMixin",
		reflect.TypeOf((*CfnStudioPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnStudioPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_nimblestudio.CfnStudioPropsMixin.StudioEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnStudioPropsMixin_StudioEncryptionConfigurationProperty)(nil)).Elem(),
	)
}

package mixinsawsevs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentMixinProps",
		reflect.TypeOf((*CfnEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin",
		reflect.TypeOf((*CfnEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.CheckProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_CheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.ConnectivityInfoProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_ConnectivityInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.HostInfoForCreateProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_HostInfoForCreateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.InitialVlanInfoProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_InitialVlanInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.InitialVlansProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_InitialVlansProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.LicenseInfoProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_LicenseInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.ServiceAccessSecurityGroupsProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_ServiceAccessSecurityGroupsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_evs.mixins.CfnEnvironmentPropsMixin.VcfHostnamesProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_VcfHostnamesProperty)(nil)).Elem(),
	)
}

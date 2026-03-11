package awsevs

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentMixinProps",
		reflect.TypeOf((*CfnEnvironmentMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin",
		reflect.TypeOf((*CfnEnvironmentPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnEnvironmentPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.CheckProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_CheckProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.ConnectivityInfoProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_ConnectivityInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.HostInfoForCreateProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_HostInfoForCreateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.InitialVlanInfoProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_InitialVlanInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.InitialVlansProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_InitialVlansProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.LicenseInfoProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_LicenseInfoProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.SecretProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_SecretProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.ServiceAccessSecurityGroupsProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_ServiceAccessSecurityGroupsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_evs.CfnEnvironmentPropsMixin.VcfHostnamesProperty",
		reflect.TypeOf((*CfnEnvironmentPropsMixin_VcfHostnamesProperty)(nil)).Elem(),
	)
}

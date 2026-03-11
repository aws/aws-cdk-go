package awsapprunner

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnAutoScalingConfigurationMixinProps",
		reflect.TypeOf((*CfnAutoScalingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnAutoScalingConfigurationPropsMixin",
		reflect.TypeOf((*CfnAutoScalingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutoScalingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationMixinProps",
		reflect.TypeOf((*CfnObservabilityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationPropsMixin",
		reflect.TypeOf((*CfnObservabilityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnObservabilityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnObservabilityConfigurationPropsMixin.TraceConfigurationProperty",
		reflect.TypeOf((*CfnObservabilityConfigurationPropsMixin_TraceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.CodeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.CodeConfigurationValuesProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CodeConfigurationValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.EgressConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EgressConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.HealthCheckConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_HealthCheckConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.ImageConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ImageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.ImageRepositoryProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ImageRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.IngressConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_IngressConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.InstanceConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_InstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnServicePropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.ServiceObservabilityConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceObservabilityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.SourceCodeVersionProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SourceCodeVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnServicePropsMixin.SourceConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnVpcConnectorMixinProps",
		reflect.TypeOf((*CfnVpcConnectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnVpcConnectorPropsMixin",
		reflect.TypeOf((*CfnVpcConnectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcConnectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnVpcIngressConnectionMixinProps",
		reflect.TypeOf((*CfnVpcIngressConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnVpcIngressConnectionPropsMixin",
		reflect.TypeOf((*CfnVpcIngressConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcIngressConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_apprunner.CfnVpcIngressConnectionPropsMixin.IngressVpcConfigurationProperty",
		reflect.TypeOf((*CfnVpcIngressConnectionPropsMixin_IngressVpcConfigurationProperty)(nil)).Elem(),
	)
}

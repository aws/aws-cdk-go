package previewawsapprunnermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnAutoScalingConfigurationMixinProps",
		reflect.TypeOf((*CfnAutoScalingConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnAutoScalingConfigurationPropsMixin",
		reflect.TypeOf((*CfnAutoScalingConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutoScalingConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnObservabilityConfigurationMixinProps",
		reflect.TypeOf((*CfnObservabilityConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnObservabilityConfigurationPropsMixin",
		reflect.TypeOf((*CfnObservabilityConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnObservabilityConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnObservabilityConfigurationPropsMixin.TraceConfigurationProperty",
		reflect.TypeOf((*CfnObservabilityConfigurationPropsMixin_TraceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServiceMixinProps",
		reflect.TypeOf((*CfnServiceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin",
		reflect.TypeOf((*CfnServicePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnServicePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.AuthenticationConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_AuthenticationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.CodeConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CodeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.CodeConfigurationValuesProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CodeConfigurationValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.CodeRepositoryProperty",
		reflect.TypeOf((*CfnServicePropsMixin_CodeRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.EgressConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EgressConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.EncryptionConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_EncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.HealthCheckConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_HealthCheckConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.ImageConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ImageConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.ImageRepositoryProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ImageRepositoryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.IngressConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_IngressConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.InstanceConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_InstanceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.KeyValuePairProperty",
		reflect.TypeOf((*CfnServicePropsMixin_KeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.ServiceObservabilityConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_ServiceObservabilityConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.SourceCodeVersionProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SourceCodeVersionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnServicePropsMixin.SourceConfigurationProperty",
		reflect.TypeOf((*CfnServicePropsMixin_SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnVpcConnectorMixinProps",
		reflect.TypeOf((*CfnVpcConnectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnVpcConnectorPropsMixin",
		reflect.TypeOf((*CfnVpcConnectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcConnectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnVpcIngressConnectionMixinProps",
		reflect.TypeOf((*CfnVpcIngressConnectionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnVpcIngressConnectionPropsMixin",
		reflect.TypeOf((*CfnVpcIngressConnectionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVpcIngressConnectionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_apprunner.mixins.CfnVpcIngressConnectionPropsMixin.IngressVpcConfigurationProperty",
		reflect.TypeOf((*CfnVpcIngressConnectionPropsMixin_IngressVpcConfigurationProperty)(nil)).Elem(),
	)
}

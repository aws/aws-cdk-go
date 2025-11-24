package mixinsawsiotcoredeviceadvisor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotcoredeviceadvisor.mixins.CfnSuiteDefinitionMixinProps",
		reflect.TypeOf((*CfnSuiteDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_iotcoredeviceadvisor.mixins.CfnSuiteDefinitionPropsMixin",
		reflect.TypeOf((*CfnSuiteDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSuiteDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotcoredeviceadvisor.mixins.CfnSuiteDefinitionPropsMixin.DeviceUnderTestProperty",
		reflect.TypeOf((*CfnSuiteDefinitionPropsMixin_DeviceUnderTestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_iotcoredeviceadvisor.mixins.CfnSuiteDefinitionPropsMixin.SuiteDefinitionConfigurationProperty",
		reflect.TypeOf((*CfnSuiteDefinitionPropsMixin_SuiteDefinitionConfigurationProperty)(nil)).Elem(),
	)
}

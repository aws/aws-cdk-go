package awsiotcoredeviceadvisor

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotcoredeviceadvisor.CfnSuiteDefinitionMixinProps",
		reflect.TypeOf((*CfnSuiteDefinitionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_iotcoredeviceadvisor.CfnSuiteDefinitionPropsMixin",
		reflect.TypeOf((*CfnSuiteDefinitionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSuiteDefinitionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotcoredeviceadvisor.CfnSuiteDefinitionPropsMixin.DeviceUnderTestProperty",
		reflect.TypeOf((*CfnSuiteDefinitionPropsMixin_DeviceUnderTestProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_iotcoredeviceadvisor.CfnSuiteDefinitionPropsMixin.SuiteDefinitionConfigurationProperty",
		reflect.TypeOf((*CfnSuiteDefinitionPropsMixin_SuiteDefinitionConfigurationProperty)(nil)).Elem(),
	)
}

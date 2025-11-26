package previewawsrummixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorMixinProps",
		reflect.TypeOf((*CfnAppMonitorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin",
		reflect.TypeOf((*CfnAppMonitorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppMonitorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.AppMonitorConfigurationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_AppMonitorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.CustomEventsProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_CustomEventsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.DeobfuscationConfigurationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_DeobfuscationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.JavaScriptSourceMapsProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_JavaScriptSourceMapsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.MetricDefinitionProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_MetricDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.MetricDestinationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_MetricDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_rum.mixins.CfnAppMonitorPropsMixin.ResourcePolicyProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_ResourcePolicyProperty)(nil)).Elem(),
	)
}

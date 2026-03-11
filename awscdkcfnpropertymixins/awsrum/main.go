package awsrum

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorMixinProps",
		reflect.TypeOf((*CfnAppMonitorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin",
		reflect.TypeOf((*CfnAppMonitorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAppMonitorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.AppMonitorConfigurationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_AppMonitorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.CustomEventsProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_CustomEventsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.DeobfuscationConfigurationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_DeobfuscationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.JavaScriptSourceMapsProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_JavaScriptSourceMapsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.MetricDefinitionProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_MetricDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.MetricDestinationProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_MetricDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_rum.CfnAppMonitorPropsMixin.ResourcePolicyProperty",
		reflect.TypeOf((*CfnAppMonitorPropsMixin_ResourcePolicyProperty)(nil)).Elem(),
	)
}

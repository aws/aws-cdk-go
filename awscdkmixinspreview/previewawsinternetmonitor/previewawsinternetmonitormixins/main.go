package previewawsinternetmonitormixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorMixinProps",
		reflect.TypeOf((*CfnMonitorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin",
		reflect.TypeOf((*CfnMonitorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMonitorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin.HealthEventsConfigProperty",
		reflect.TypeOf((*CfnMonitorPropsMixin_HealthEventsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin.InternetMeasurementsLogDeliveryProperty",
		reflect.TypeOf((*CfnMonitorPropsMixin_InternetMeasurementsLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin.LocalHealthEventsConfigProperty",
		reflect.TypeOf((*CfnMonitorPropsMixin_LocalHealthEventsConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_internetmonitor.mixins.CfnMonitorPropsMixin.S3ConfigProperty",
		reflect.TypeOf((*CfnMonitorPropsMixin_S3ConfigProperty)(nil)).Elem(),
	)
}

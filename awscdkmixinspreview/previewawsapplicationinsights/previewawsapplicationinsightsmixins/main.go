package previewawsapplicationinsightsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.AlarmMetricProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AlarmMetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.AlarmProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AlarmProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.ComponentConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ComponentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.ComponentMonitoringSettingProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ComponentMonitoringSettingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.ConfigurationDetailsProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ConfigurationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.CustomComponentProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CustomComponentProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.HAClusterPrometheusExporterProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_HAClusterPrometheusExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.HANAPrometheusExporterProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_HANAPrometheusExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.JMXPrometheusExporterProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_JMXPrometheusExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.LogPatternProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_LogPatternProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.LogPatternSetProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_LogPatternSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.LogProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_LogProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.NetWeaverPrometheusExporterProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_NetWeaverPrometheusExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.ProcessProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ProcessProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.SQLServerPrometheusExporterProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_SQLServerPrometheusExporterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.SubComponentConfigurationDetailsProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_SubComponentConfigurationDetailsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.SubComponentTypeConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_SubComponentTypeConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_applicationinsights.mixins.CfnApplicationPropsMixin.WindowsEventProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_WindowsEventProperty)(nil)).Elem(),
	)
}

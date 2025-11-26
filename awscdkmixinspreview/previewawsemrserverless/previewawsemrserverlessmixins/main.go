package previewawsemrserverlessmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin",
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
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.AutoStartConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AutoStartConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.AutoStopConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_AutoStopConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.CloudWatchLoggingConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CloudWatchLoggingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.ConfigurationObjectProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ConfigurationObjectProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.IdentityCenterConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_IdentityCenterConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.ImageConfigurationInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ImageConfigurationInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.InitialCapacityConfigKeyValuePairProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InitialCapacityConfigKeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.InitialCapacityConfigProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InitialCapacityConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.InteractiveConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InteractiveConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.LogTypeMapKeyValuePairProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_LogTypeMapKeyValuePairProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.ManagedPersistenceMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_ManagedPersistenceMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.MaximumAllowedResourcesProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MaximumAllowedResourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.NetworkConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_NetworkConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.PrometheusMonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_PrometheusMonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.S3MonitoringConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_S3MonitoringConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.SchedulerConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_SchedulerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.WorkerConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_WorkerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_emrserverless.mixins.CfnApplicationPropsMixin.WorkerTypeSpecificationInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_WorkerTypeSpecificationInputProperty)(nil)).Elem(),
	)
}

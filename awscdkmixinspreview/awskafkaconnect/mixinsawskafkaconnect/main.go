package mixinsawskafkaconnect

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorMixinProps",
		reflect.TypeOf((*CfnConnectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin",
		reflect.TypeOf((*CfnConnectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnConnectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.ApacheKafkaClusterProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_ApacheKafkaClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.AutoScalingProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_AutoScalingProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.CapacityProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_CapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.CloudWatchLogsLogDeliveryProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_CloudWatchLogsLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.CustomPluginProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_CustomPluginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.FirehoseLogDeliveryProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_FirehoseLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.KafkaClusterClientAuthenticationProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_KafkaClusterClientAuthenticationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.KafkaClusterEncryptionInTransitProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_KafkaClusterEncryptionInTransitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.KafkaClusterProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_KafkaClusterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.LogDeliveryProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_LogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.PluginProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_PluginProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.ProvisionedCapacityProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_ProvisionedCapacityProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.S3LogDeliveryProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_S3LogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.ScaleInPolicyProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_ScaleInPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.ScaleOutPolicyProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_ScaleOutPolicyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.VpcProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_VpcProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.WorkerConfigurationProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_WorkerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnConnectorPropsMixin.WorkerLogDeliveryProperty",
		reflect.TypeOf((*CfnConnectorPropsMixin_WorkerLogDeliveryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnCustomPluginMixinProps",
		reflect.TypeOf((*CfnCustomPluginMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnCustomPluginPropsMixin",
		reflect.TypeOf((*CfnCustomPluginPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCustomPluginPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnCustomPluginPropsMixin.CustomPluginFileDescriptionProperty",
		reflect.TypeOf((*CfnCustomPluginPropsMixin_CustomPluginFileDescriptionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnCustomPluginPropsMixin.CustomPluginLocationProperty",
		reflect.TypeOf((*CfnCustomPluginPropsMixin_CustomPluginLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnCustomPluginPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnCustomPluginPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnWorkerConfigurationMixinProps",
		reflect.TypeOf((*CfnWorkerConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_kafkaconnect.mixins.CfnWorkerConfigurationPropsMixin",
		reflect.TypeOf((*CfnWorkerConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnWorkerConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}

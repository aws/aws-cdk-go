package previewawsobservabilityadminmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRuleMixinProps",
		reflect.TypeOf((*CfnOrganizationCentralizationRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationCentralizationRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.CentralizationRuleDestinationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.CentralizationRuleProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.CentralizationRuleSourceProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.DestinationLogsConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_DestinationLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.LogsBackupConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_LogsBackupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.LogsEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_LogsEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationCentralizationRulePropsMixin.SourceLogsConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_SourceLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRuleMixinProps",
		reflect.TypeOf((*CfnOrganizationTelemetryRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationTelemetryRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin.TelemetryDestinationConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_TelemetryDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin.TelemetryRuleProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_TelemetryRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnOrganizationTelemetryRulePropsMixin.VPCFlowLogParametersProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_VPCFlowLogParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationMixinProps",
		reflect.TypeOf((*CfnS3TableIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin",
		reflect.TypeOf((*CfnS3TableIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnS3TableIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin.EncryptionConfigProperty",
		reflect.TypeOf((*CfnS3TableIntegrationPropsMixin_EncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnS3TableIntegrationPropsMixin.LogSourceProperty",
		reflect.TypeOf((*CfnS3TableIntegrationPropsMixin_LogSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesMixinProps",
		reflect.TypeOf((*CfnTelemetryPipelinesMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTelemetryPipelinesPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin.TelemetryPipelineConfigurationProperty",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin_TelemetryPipelineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin.TelemetryPipelineProperty",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin_TelemetryPipelineProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryPipelinesPropsMixin.TelemetryPipelineStatusReasonProperty",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin_TelemetryPipelineStatusReasonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRuleMixinProps",
		reflect.TypeOf((*CfnTelemetryRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTelemetryRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin.TelemetryDestinationConfigurationProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_TelemetryDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin.TelemetryRuleProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_TelemetryRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_observabilityadmin.mixins.CfnTelemetryRulePropsMixin.VPCFlowLogParametersProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_VPCFlowLogParametersProperty)(nil)).Elem(),
	)
}

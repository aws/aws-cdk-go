package awsobservabilityadmin

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRuleMixinProps",
		reflect.TypeOf((*CfnOrganizationCentralizationRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationCentralizationRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.CentralizationRuleDestinationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleDestinationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.CentralizationRuleProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.CentralizationRuleSourceProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_CentralizationRuleSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.DestinationLogsConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_DestinationLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.LogGroupNameConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_LogGroupNameConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.LogsBackupConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_LogsBackupConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.LogsEncryptionConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_LogsEncryptionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationCentralizationRulePropsMixin.SourceLogsConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationCentralizationRulePropsMixin_SourceLogsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRuleMixinProps",
		reflect.TypeOf((*CfnOrganizationTelemetryRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnOrganizationTelemetryRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.ActionConditionProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_ActionConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.AdvancedEventSelectorProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_AdvancedEventSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.AdvancedFieldSelectorProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_AdvancedFieldSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.CloudtrailParametersProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_CloudtrailParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.ELBLoadBalancerLoggingParametersProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_ELBLoadBalancerLoggingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.FieldToMatchProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.FilterProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.LabelNameConditionProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_LabelNameConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.LoggingFilterProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_LoggingFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.SingleHeaderProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_SingleHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.TelemetryDestinationConfigurationProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_TelemetryDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.TelemetryRuleProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_TelemetryRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.VPCFlowLogParametersProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_VPCFlowLogParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnOrganizationTelemetryRulePropsMixin.WAFLoggingParametersProperty",
		reflect.TypeOf((*CfnOrganizationTelemetryRulePropsMixin_WAFLoggingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnS3TableIntegrationMixinProps",
		reflect.TypeOf((*CfnS3TableIntegrationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnS3TableIntegrationPropsMixin",
		reflect.TypeOf((*CfnS3TableIntegrationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnS3TableIntegrationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnS3TableIntegrationPropsMixin.EncryptionConfigProperty",
		reflect.TypeOf((*CfnS3TableIntegrationPropsMixin_EncryptionConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnS3TableIntegrationPropsMixin.LogSourceProperty",
		reflect.TypeOf((*CfnS3TableIntegrationPropsMixin_LogSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryPipelinesMixinProps",
		reflect.TypeOf((*CfnTelemetryPipelinesMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryPipelinesPropsMixin",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTelemetryPipelinesPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryPipelinesPropsMixin.TelemetryPipelineConfigurationProperty",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin_TelemetryPipelineConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryPipelinesPropsMixin.TelemetryPipelineProperty",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin_TelemetryPipelineProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryPipelinesPropsMixin.TelemetryPipelineStatusReasonProperty",
		reflect.TypeOf((*CfnTelemetryPipelinesPropsMixin_TelemetryPipelineStatusReasonProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRuleMixinProps",
		reflect.TypeOf((*CfnTelemetryRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTelemetryRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.ActionConditionProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_ActionConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.AdvancedEventSelectorProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_AdvancedEventSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.AdvancedFieldSelectorProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_AdvancedFieldSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.CloudtrailParametersProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_CloudtrailParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.ConditionProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_ConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.ELBLoadBalancerLoggingParametersProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_ELBLoadBalancerLoggingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.FieldToMatchProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_FieldToMatchProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.FilterProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.LabelNameConditionProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_LabelNameConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.LogDeliveryParametersProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_LogDeliveryParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.LoggingFilterProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_LoggingFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.SingleHeaderProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_SingleHeaderProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.TelemetryDestinationConfigurationProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_TelemetryDestinationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.TelemetryRuleProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_TelemetryRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.VPCFlowLogParametersProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_VPCFlowLogParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_observabilityadmin.CfnTelemetryRulePropsMixin.WAFLoggingParametersProperty",
		reflect.TypeOf((*CfnTelemetryRulePropsMixin_WAFLoggingParametersProperty)(nil)).Elem(),
	)
}

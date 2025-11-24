package mixinsawsssmincidents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnReplicationSetMixinProps",
		reflect.TypeOf((*CfnReplicationSetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnReplicationSetPropsMixin",
		reflect.TypeOf((*CfnReplicationSetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnReplicationSetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnReplicationSetPropsMixin.RegionConfigurationProperty",
		reflect.TypeOf((*CfnReplicationSetPropsMixin_RegionConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnReplicationSetPropsMixin.ReplicationRegionProperty",
		reflect.TypeOf((*CfnReplicationSetPropsMixin_ReplicationRegionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanMixinProps",
		reflect.TypeOf((*CfnResponsePlanMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin",
		reflect.TypeOf((*CfnResponsePlanPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnResponsePlanPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.ChatChannelProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_ChatChannelProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.DynamicSsmParameterProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_DynamicSsmParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.DynamicSsmParameterValueProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_DynamicSsmParameterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.IncidentTemplateProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_IncidentTemplateProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.IntegrationProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_IntegrationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.NotificationTargetItemProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_NotificationTargetItemProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.PagerDutyConfigurationProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_PagerDutyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.PagerDutyIncidentConfigurationProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_PagerDutyIncidentConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.SsmAutomationProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_SsmAutomationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_ssmincidents.mixins.CfnResponsePlanPropsMixin.SsmParameterProperty",
		reflect.TypeOf((*CfnResponsePlanPropsMixin_SsmParameterProperty)(nil)).Elem(),
	)
}

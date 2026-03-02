package previewawscomputeoptimizermixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRuleMixinProps",
		reflect.TypeOf((*CfnAutomationRuleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin",
		reflect.TypeOf((*CfnAutomationRulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAutomationRulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.CriteriaProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_CriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.DoubleCriteriaConditionProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_DoubleCriteriaConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.IntegerCriteriaConditionProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_IntegerCriteriaConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.OrganizationConfigurationProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_OrganizationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.ResourceTagsCriteriaConditionProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_ResourceTagsCriteriaConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.ScheduleProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_ScheduleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_computeoptimizer.mixins.CfnAutomationRulePropsMixin.StringCriteriaConditionProperty",
		reflect.TypeOf((*CfnAutomationRulePropsMixin_StringCriteriaConditionProperty)(nil)).Elem(),
	)
}

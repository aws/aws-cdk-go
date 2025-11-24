package mixinsawsbudgets

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetMixinProps",
		reflect.TypeOf((*CfnBudgetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin",
		reflect.TypeOf((*CfnBudgetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBudgetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.AutoAdjustDataProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_AutoAdjustDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.BudgetDataProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_BudgetDataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.CostCategoryValuesProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_CostCategoryValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.CostTypesProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_CostTypesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.ExpressionDimensionValuesProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_ExpressionDimensionValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.ExpressionProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_ExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.HistoricalOptionsProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_HistoricalOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.NotificationProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_NotificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.NotificationWithSubscribersProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_NotificationWithSubscribersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.SpendProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_SpendProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.SubscriberProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_SubscriberProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.TagValuesProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_TagValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetPropsMixin.TimePeriodProperty",
		reflect.TypeOf((*CfnBudgetPropsMixin_TimePeriodProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionMixinProps",
		reflect.TypeOf((*CfnBudgetsActionMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnBudgetsActionPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.ActionThresholdProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_ActionThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.DefinitionProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_DefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.IamActionDefinitionProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_IamActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.ResourceTagProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_ResourceTagProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.ScpActionDefinitionProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_ScpActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.SsmActionDefinitionProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_SsmActionDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_budgets.mixins.CfnBudgetsActionPropsMixin.SubscriberProperty",
		reflect.TypeOf((*CfnBudgetsActionPropsMixin_SubscriberProperty)(nil)).Elem(),
	)
}

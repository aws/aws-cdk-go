package awsbcm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardMixinProps",
		reflect.TypeOf((*CfnDashboardMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin",
		reflect.TypeOf((*CfnDashboardPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDashboardPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.CostAndUsageExpressionProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_CostAndUsageExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.CostAndUsageQueryProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_CostAndUsageQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.CostCategoryValuesProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_CostCategoryValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.DateTimeRangeProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_DateTimeRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.DateTimeValueProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_DateTimeValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.DimensionValuesProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_DimensionValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.DisplayConfigProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_DisplayConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.ExpressionProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_ExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.GraphDisplayConfigProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_GraphDisplayConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.GroupDefinitionProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_GroupDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.QueryParametersProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_QueryParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.ReservationCoverageQueryProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_ReservationCoverageQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.ReservationUtilizationQueryProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_ReservationUtilizationQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.SavingsPlansCoverageQueryProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_SavingsPlansCoverageQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.SavingsPlansUtilizationQueryProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_SavingsPlansUtilizationQueryProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.TagValuesProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_TagValuesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.WidgetConfigProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_WidgetConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_bcm.CfnDashboardPropsMixin.WidgetProperty",
		reflect.TypeOf((*CfnDashboardPropsMixin_WidgetProperty)(nil)).Elem(),
	)
}

package awsdatabrew

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetMixinProps",
		reflect.TypeOf((*CfnDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin",
		reflect.TypeOf((*CfnDatasetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatasetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.CsvOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_CsvOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.DataCatalogInputDefinitionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DataCatalogInputDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.DatabaseInputDefinitionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatabaseInputDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.DatasetParameterProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatasetParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.DatetimeOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatetimeOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.ExcelOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_ExcelOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.FilesLimitProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilesLimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.FilterExpressionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilterExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.FilterValueProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.FormatOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.InputProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.JsonOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_JsonOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.MetadataProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_MetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.PathOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_PathOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.PathParameterProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_PathParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnDatasetPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobMixinProps",
		reflect.TypeOf((*CfnJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin",
		reflect.TypeOf((*CfnJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.AllowedStatisticsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_AllowedStatisticsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.ColumnSelectorProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ColumnSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.ColumnStatisticsConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ColumnStatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.CsvOutputOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_CsvOutputOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.DataCatalogOutputProperty",
		reflect.TypeOf((*CfnJobPropsMixin_DataCatalogOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.DatabaseOutputProperty",
		reflect.TypeOf((*CfnJobPropsMixin_DatabaseOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.DatabaseTableOutputOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_DatabaseTableOutputOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.EntityDetectorConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_EntityDetectorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.JobSampleProperty",
		reflect.TypeOf((*CfnJobPropsMixin_JobSampleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.OutputFormatOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_OutputFormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.OutputLocationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_OutputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.OutputProperty",
		reflect.TypeOf((*CfnJobPropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.ProfileConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ProfileConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.RecipeProperty",
		reflect.TypeOf((*CfnJobPropsMixin_RecipeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.S3TableOutputOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_S3TableOutputOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.StatisticOverrideProperty",
		reflect.TypeOf((*CfnJobPropsMixin_StatisticOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.StatisticsConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_StatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnJobPropsMixin.ValidationConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ValidationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnProjectPropsMixin.SampleProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_SampleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipeMixinProps",
		reflect.TypeOf((*CfnRecipeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin",
		reflect.TypeOf((*CfnRecipePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecipePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin.ActionProperty",
		reflect.TypeOf((*CfnRecipePropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin.ConditionExpressionProperty",
		reflect.TypeOf((*CfnRecipePropsMixin_ConditionExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRecipePropsMixin.RecipeStepProperty",
		reflect.TypeOf((*CfnRecipePropsMixin_RecipeStepProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRulesetMixinProps",
		reflect.TypeOf((*CfnRulesetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRulesetPropsMixin",
		reflect.TypeOf((*CfnRulesetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRulesetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRulesetPropsMixin.ColumnSelectorProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_ColumnSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRulesetPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRulesetPropsMixin.SubstitutionValueProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_SubstitutionValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnRulesetPropsMixin.ThresholdProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_ThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnScheduleMixinProps",
		reflect.TypeOf((*CfnScheduleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_databrew.CfnSchedulePropsMixin",
		reflect.TypeOf((*CfnSchedulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchedulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}

package mixinsawsdatabrew

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetMixinProps",
		reflect.TypeOf((*CfnDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin",
		reflect.TypeOf((*CfnDatasetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDatasetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.CsvOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_CsvOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.DataCatalogInputDefinitionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DataCatalogInputDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.DatabaseInputDefinitionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatabaseInputDefinitionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.DatasetParameterProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatasetParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.DatetimeOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DatetimeOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.ExcelOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_ExcelOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.FilesLimitProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilesLimitProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.FilterExpressionProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilterExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.FilterValueProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FilterValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.FormatOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_FormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.InputProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.JsonOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_JsonOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.MetadataProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_MetadataProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.PathOptionsProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_PathOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.PathParameterProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_PathParameterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnDatasetPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobMixinProps",
		reflect.TypeOf((*CfnJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin",
		reflect.TypeOf((*CfnJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.AllowedStatisticsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_AllowedStatisticsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.ColumnSelectorProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ColumnSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.ColumnStatisticsConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ColumnStatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.CsvOutputOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_CsvOutputOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.DataCatalogOutputProperty",
		reflect.TypeOf((*CfnJobPropsMixin_DataCatalogOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.DatabaseOutputProperty",
		reflect.TypeOf((*CfnJobPropsMixin_DatabaseOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.DatabaseTableOutputOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_DatabaseTableOutputOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.EntityDetectorConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_EntityDetectorConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.JobSampleProperty",
		reflect.TypeOf((*CfnJobPropsMixin_JobSampleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.OutputFormatOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_OutputFormatOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.OutputLocationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_OutputLocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.OutputProperty",
		reflect.TypeOf((*CfnJobPropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.ProfileConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ProfileConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.RecipeProperty",
		reflect.TypeOf((*CfnJobPropsMixin_RecipeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.S3LocationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_S3LocationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.S3TableOutputOptionsProperty",
		reflect.TypeOf((*CfnJobPropsMixin_S3TableOutputOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.StatisticOverrideProperty",
		reflect.TypeOf((*CfnJobPropsMixin_StatisticOverrideProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.StatisticsConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_StatisticsConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnJobPropsMixin.ValidationConfigurationProperty",
		reflect.TypeOf((*CfnJobPropsMixin_ValidationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnProjectMixinProps",
		reflect.TypeOf((*CfnProjectMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnProjectPropsMixin",
		reflect.TypeOf((*CfnProjectPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnProjectPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnProjectPropsMixin.SampleProperty",
		reflect.TypeOf((*CfnProjectPropsMixin_SampleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRecipeMixinProps",
		reflect.TypeOf((*CfnRecipeMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRecipePropsMixin",
		reflect.TypeOf((*CfnRecipePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRecipePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRecipePropsMixin.ActionProperty",
		reflect.TypeOf((*CfnRecipePropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRecipePropsMixin.ConditionExpressionProperty",
		reflect.TypeOf((*CfnRecipePropsMixin_ConditionExpressionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRecipePropsMixin.RecipeStepProperty",
		reflect.TypeOf((*CfnRecipePropsMixin_RecipeStepProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetMixinProps",
		reflect.TypeOf((*CfnRulesetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin",
		reflect.TypeOf((*CfnRulesetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnRulesetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin.ColumnSelectorProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_ColumnSelectorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin.SubstitutionValueProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_SubstitutionValueProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnRulesetPropsMixin.ThresholdProperty",
		reflect.TypeOf((*CfnRulesetPropsMixin_ThresholdProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnScheduleMixinProps",
		reflect.TypeOf((*CfnScheduleMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_databrew.mixins.CfnSchedulePropsMixin",
		reflect.TypeOf((*CfnSchedulePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnSchedulePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
}

package previewawscleanroomsmlmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetMixinProps",
		reflect.TypeOf((*CfnTrainingDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrainingDatasetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin.ColumnSchemaProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_ColumnSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin.DataSourceProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin.DatasetInputConfigProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_DatasetInputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin.DatasetProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_DatasetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_cleanroomsml.mixins.CfnTrainingDatasetPropsMixin.GlueDataSourceProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_GlueDataSourceProperty)(nil)).Elem(),
	)
}

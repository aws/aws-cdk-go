package awscleanroomsml

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetMixinProps",
		reflect.TypeOf((*CfnTrainingDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetPropsMixin",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTrainingDatasetPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetPropsMixin.ColumnSchemaProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_ColumnSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetPropsMixin.DataSourceProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_DataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetPropsMixin.DatasetInputConfigProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_DatasetInputConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetPropsMixin.DatasetProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_DatasetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_cleanroomsml.CfnTrainingDatasetPropsMixin.GlueDataSourceProperty",
		reflect.TypeOf((*CfnTrainingDatasetPropsMixin_GlueDataSourceProperty)(nil)).Elem(),
	)
}

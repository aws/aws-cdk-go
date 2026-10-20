package awsscn

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowMixinProps",
		reflect.TypeOf((*CfnDataIntegrationFlowMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnDataIntegrationFlowPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DataIntegrationFlowSourceProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DataIntegrationFlowSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DataIntegrationFlowTargetProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DataIntegrationFlowTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DataIntegrationFlowTransformationProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DataIntegrationFlowTransformationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DatasetOptionsProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DatasetOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DatasetSourceConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DatasetSourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DatasetTargetConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DatasetTargetConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.DedupeStrategyProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_DedupeStrategyProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.FieldPriorityDedupeFieldProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_FieldPriorityDedupeFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.FieldPriorityDedupeStrategyConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_FieldPriorityDedupeStrategyConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.S3OptionsProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_S3OptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.S3SourceConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_S3SourceConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDataIntegrationFlowPropsMixin.SqlTransformationConfigurationProperty",
		reflect.TypeOf((*CfnDataIntegrationFlowPropsMixin_SqlTransformationConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetMixinProps",
		reflect.TypeOf((*CfnDatasetMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin",
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
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin.DataLakeDatasetPartitionFieldProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DataLakeDatasetPartitionFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin.DataLakeDatasetPrimaryKeyFieldProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DataLakeDatasetPrimaryKeyFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin.DataLakeDatasetSchemaFieldProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_DataLakeDatasetSchemaFieldProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin.PartitionSpecProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_PartitionSpecProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin.SchemaProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_SchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnDatasetPropsMixin.TransformProperty",
		reflect.TypeOf((*CfnDatasetPropsMixin_TransformProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnInstanceMixinProps",
		reflect.TypeOf((*CfnInstanceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnInstancePropsMixin",
		reflect.TypeOf((*CfnInstancePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnInstancePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnNamespaceMixinProps",
		reflect.TypeOf((*CfnNamespaceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_scn.CfnNamespacePropsMixin",
		reflect.TypeOf((*CfnNamespacePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnNamespacePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}

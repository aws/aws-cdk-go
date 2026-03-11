package awskinesisanalytics

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationMixinProps",
		reflect.TypeOf((*CfnApplicationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputMixinProps",
		reflect.TypeOf((*CfnApplicationOutputMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputPropsMixin",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationOutputPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputPropsMixin.DestinationSchemaProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_DestinationSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputPropsMixin.KinesisFirehoseOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_KinesisFirehoseOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputPropsMixin.KinesisStreamsOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_KinesisStreamsOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputPropsMixin.LambdaOutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_LambdaOutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationOutputPropsMixin.OutputProperty",
		reflect.TypeOf((*CfnApplicationOutputPropsMixin_OutputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin",
		reflect.TypeOf((*CfnApplicationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.InputLambdaProcessorProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputLambdaProcessorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.InputParallelismProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputParallelismProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.InputProcessingConfigurationProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputProcessingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.InputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.InputSchemaProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_InputSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.KinesisFirehoseInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_KinesisFirehoseInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.KinesisStreamsInputProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_KinesisStreamsInputProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationPropsMixin.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationPropsMixin_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourceMixinProps",
		reflect.TypeOf((*CfnApplicationReferenceDataSourceMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnApplicationReferenceDataSourcePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.CSVMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_CSVMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.JSONMappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_JSONMappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.MappingParametersProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_MappingParametersProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.RecordColumnProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_RecordColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.RecordFormatProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_RecordFormatProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_ReferenceDataSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.ReferenceSchemaProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_ReferenceSchemaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_kinesisanalytics.CfnApplicationReferenceDataSourcePropsMixin.S3ReferenceDataSourceProperty",
		reflect.TypeOf((*CfnApplicationReferenceDataSourcePropsMixin_S3ReferenceDataSourceProperty)(nil)).Elem(),
	)
}

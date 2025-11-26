package previewawslookoutmetricsmixins

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertMixinProps",
		reflect.TypeOf((*CfnAlertMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin",
		reflect.TypeOf((*CfnAlertPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlertPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin.LambdaConfigurationProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_LambdaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAlertPropsMixin.SNSConfigurationProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_SNSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.AnomalyDetectorConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_AnomalyDetectorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.AppFlowConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_AppFlowConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.CloudwatchConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_CloudwatchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.CsvFormatDescriptorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_CsvFormatDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.FileFormatDescriptorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_FileFormatDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.JsonFormatDescriptorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_JsonFormatDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.MetricSetProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.MetricSourceProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.RDSSourceConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RDSSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.RedshiftSourceConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RedshiftSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.S3SourceConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_S3SourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.TimestampColumnProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_TimestampColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_lookoutmetrics.mixins.CfnAnomalyDetectorPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
}

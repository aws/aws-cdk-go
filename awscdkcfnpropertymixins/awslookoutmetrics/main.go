package awslookoutmetrics

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAlertMixinProps",
		reflect.TypeOf((*CfnAlertMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAlertPropsMixin",
		reflect.TypeOf((*CfnAlertPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAlertPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAlertPropsMixin.ActionProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_ActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAlertPropsMixin.LambdaConfigurationProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_LambdaConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAlertPropsMixin.SNSConfigurationProperty",
		reflect.TypeOf((*CfnAlertPropsMixin_SNSConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorMixinProps",
		reflect.TypeOf((*CfnAnomalyDetectorMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnomalyDetectorPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.AnomalyDetectorConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_AnomalyDetectorConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.AppFlowConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_AppFlowConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.CloudwatchConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_CloudwatchConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.CsvFormatDescriptorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_CsvFormatDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.FileFormatDescriptorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_FileFormatDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.JsonFormatDescriptorProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_JsonFormatDescriptorProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.MetricProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.MetricSetProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricSetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.MetricSourceProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_MetricSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.RDSSourceConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RDSSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.RedshiftSourceConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_RedshiftSourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.S3SourceConfigProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_S3SourceConfigProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.TimestampColumnProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_TimestampColumnProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_lookoutmetrics.CfnAnomalyDetectorPropsMixin.VpcConfigurationProperty",
		reflect.TypeOf((*CfnAnomalyDetectorPropsMixin_VpcConfigurationProperty)(nil)).Elem(),
	)
}

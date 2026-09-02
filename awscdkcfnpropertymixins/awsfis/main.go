package awsfis

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplateMixinProps",
		reflect.TypeOf((*CfnExperimentTemplateMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnExperimentTemplatePropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.CloudWatchDashboardProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_CloudWatchDashboardProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.DataSourcesProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_DataSourcesProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentReportS3ConfigurationProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentReportS3ConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateActionProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateActionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateExperimentOptionsProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateExperimentOptionsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateExperimentReportConfigurationProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateExperimentReportConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateLogConfigurationProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateLogConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateStopConditionProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateStopConditionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateTargetFilterProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateTargetFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.ExperimentTemplateTargetProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_ExperimentTemplateTargetProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnExperimentTemplatePropsMixin.OutputsProperty",
		reflect.TypeOf((*CfnExperimentTemplatePropsMixin_OutputsProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnTargetAccountConfigurationMixinProps",
		reflect.TypeOf((*CfnTargetAccountConfigurationMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_fis.CfnTargetAccountConfigurationPropsMixin",
		reflect.TypeOf((*CfnTargetAccountConfigurationPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnTargetAccountConfigurationPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}

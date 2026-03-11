package awsaccessanalyzer

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerMixinProps",
		reflect.TypeOf((*CfnAnalyzerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin",
		reflect.TypeOf((*CfnAnalyzerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnalyzerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.AnalysisRuleCriteriaProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_AnalysisRuleCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.AnalysisRuleProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_AnalysisRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.AnalyzerConfigurationProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_AnalyzerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.ArchiveRuleProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_ArchiveRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.InternalAccessAnalysisRuleCriteriaProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_InternalAccessAnalysisRuleCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.InternalAccessAnalysisRuleProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_InternalAccessAnalysisRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.InternalAccessConfigurationProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_InternalAccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_accessanalyzer.CfnAnalyzerPropsMixin.UnusedAccessConfigurationProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_UnusedAccessConfigurationProperty)(nil)).Elem(),
	)
}

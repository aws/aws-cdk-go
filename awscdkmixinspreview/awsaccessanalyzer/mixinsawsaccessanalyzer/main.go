package mixinsawsaccessanalyzer

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerMixinProps",
		reflect.TypeOf((*CfnAnalyzerMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin",
		reflect.TypeOf((*CfnAnalyzerPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnAnalyzerPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__coreMixin)
			_jsii_.InitJsiiProxy(&j.Type__coreIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.AnalysisRuleCriteriaProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_AnalysisRuleCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.AnalysisRuleProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_AnalysisRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.AnalyzerConfigurationProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_AnalyzerConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.ArchiveRuleProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_ArchiveRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.FilterProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_FilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.InternalAccessAnalysisRuleCriteriaProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_InternalAccessAnalysisRuleCriteriaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.InternalAccessAnalysisRuleProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_InternalAccessAnalysisRuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.InternalAccessConfigurationProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_InternalAccessConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_accessanalyzer.mixins.CfnAnalyzerPropsMixin.UnusedAccessConfigurationProperty",
		reflect.TypeOf((*CfnAnalyzerPropsMixin_UnusedAccessConfigurationProperty)(nil)).Elem(),
	)
}

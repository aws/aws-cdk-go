package awstranscribe

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryMixinProps",
		reflect.TypeOf((*CfnCallAnalyticsCategoryMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnCallAnalyticsCategoryPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.AbsoluteTimeRangeProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_AbsoluteTimeRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.InterruptionFilterProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_InterruptionFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.NonTalkTimeFilterProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_NonTalkTimeFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.RelativeTimeRangeProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_RelativeTimeRangeProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.RuleProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_RuleProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.SentimentFilterProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_SentimentFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnCallAnalyticsCategoryPropsMixin.TranscriptFilterProperty",
		reflect.TypeOf((*CfnCallAnalyticsCategoryPropsMixin_TranscriptFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalVocabularyMixinProps",
		reflect.TypeOf((*CfnMedicalVocabularyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalVocabularyPropsMixin",
		reflect.TypeOf((*CfnMedicalVocabularyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMedicalVocabularyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyFilterMixinProps",
		reflect.TypeOf((*CfnVocabularyFilterMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyFilterPropsMixin",
		reflect.TypeOf((*CfnVocabularyFilterPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVocabularyFilterPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyMixinProps",
		reflect.TypeOf((*CfnVocabularyMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnVocabularyPropsMixin",
		reflect.TypeOf((*CfnVocabularyPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnVocabularyPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
}

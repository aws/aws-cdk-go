package awstranscribe

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobMixinProps",
		reflect.TypeOf((*CfnMedicalTranscriptionJobMixinProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin",
		reflect.TypeOf((*CfnMedicalTranscriptionJobPropsMixin)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "applyTo", GoMethod: "ApplyTo"},
			_jsii_.MemberProperty{JsiiProperty: "props", GoGetter: "Props"},
			_jsii_.MemberProperty{JsiiProperty: "strategy", GoGetter: "Strategy"},
			_jsii_.MemberMethod{JsiiMethod: "supports", GoMethod: "Supports"},
		},
		func() interface{} {
			j := jsiiProxy_CfnMedicalTranscriptionJobPropsMixin{}
			_jsii_.InitJsiiProxy(&j.Type__awscdkMixin)
			_jsii_.InitJsiiProxy(&j.Type__constructsIMixin)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin.MediaProperty",
		reflect.TypeOf((*CfnMedicalTranscriptionJobPropsMixin_MediaProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin.MedicalTranscriptProperty",
		reflect.TypeOf((*CfnMedicalTranscriptionJobPropsMixin_MedicalTranscriptProperty)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/cfn-property-mixins.aws_transcribe.CfnMedicalTranscriptionJobPropsMixin.MedicalTranscriptionSettingProperty",
		reflect.TypeOf((*CfnMedicalTranscriptionJobPropsMixin_MedicalTranscriptionSettingProperty)(nil)).Elem(),
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
}

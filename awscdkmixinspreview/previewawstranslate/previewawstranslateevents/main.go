package previewawstranslateevents

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateParallelDataStateChange",
		reflect.TypeOf((*TranslateParallelDataStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TranslateParallelDataStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateParallelDataStateChange.TranslateParallelDataStateChangeProps",
		reflect.TypeOf((*TranslateParallelDataStateChange_TranslateParallelDataStateChangeProps)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateTextTranslationJobStateChange",
		reflect.TypeOf((*TranslateTextTranslationJobStateChange)(nil)).Elem(),
		nil, // no members
		func() interface{} {
			return &jsiiProxy_TranslateTextTranslationJobStateChange{}
		},
	)
	_jsii_.RegisterStruct(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateTextTranslationJobStateChange.TranslateTextTranslationJobStateChangeProps",
		reflect.TypeOf((*TranslateTextTranslationJobStateChange_TranslateTextTranslationJobStateChangeProps)(nil)).Elem(),
	)
}

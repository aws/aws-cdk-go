package previewawstranslateevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.translate@TranslateTextTranslationJobStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   translateTextTranslationJobStateChange := awscdkmixinspreview.Events.NewTranslateTextTranslationJobStateChange()
//
// Experimental.
type TranslateTextTranslationJobStateChange interface {
}

// The jsii proxy struct for TranslateTextTranslationJobStateChange
type jsiiProxy_TranslateTextTranslationJobStateChange struct {
	_ byte // padding
}

// Experimental.
func NewTranslateTextTranslationJobStateChange() TranslateTextTranslationJobStateChange {
	_init_.Initialize()

	j := jsiiProxy_TranslateTextTranslationJobStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateTextTranslationJobStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTranslateTextTranslationJobStateChange_Override(t TranslateTextTranslationJobStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateTextTranslationJobStateChange",
		nil, // no parameters
		t,
	)
}

// EventBridge event pattern for Translate Parallel Data State Change.
// Experimental.
func TranslateTextTranslationJobStateChange_EventPattern(options *TranslateTextTranslationJobStateChange_TranslateTextTranslationJobStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateTranslateTextTranslationJobStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateTextTranslationJobStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


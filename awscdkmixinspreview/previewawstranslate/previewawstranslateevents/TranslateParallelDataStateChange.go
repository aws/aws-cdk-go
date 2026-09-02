package previewawstranslateevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.translate@TranslateParallelDataStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   translateParallelDataStateChange := awscdkmixinspreview.Events.NewTranslateParallelDataStateChange()
//
// Experimental.
type TranslateParallelDataStateChange interface {
}

// The jsii proxy struct for TranslateParallelDataStateChange
type jsiiProxy_TranslateParallelDataStateChange struct {
	_ byte // padding
}

// Experimental.
func NewTranslateParallelDataStateChange() TranslateParallelDataStateChange {
	_init_.Initialize()

	j := jsiiProxy_TranslateParallelDataStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateParallelDataStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTranslateParallelDataStateChange_Override(t TranslateParallelDataStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateParallelDataStateChange",
		nil, // no parameters
		t,
	)
}

// EventBridge event pattern for Translate TextTranslationJob State Change.
// Experimental.
func TranslateParallelDataStateChange_EventPattern(options *TranslateParallelDataStateChange_TranslateParallelDataStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateTranslateParallelDataStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_translate.events.TranslateParallelDataStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawsb2bievents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.b2bi@TransformationCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   transformationCompleted := awscdkmixinspreview.Events.NewTransformationCompleted()
//
// Experimental.
type TransformationCompleted interface {
}

// The jsii proxy struct for TransformationCompleted
type jsiiProxy_TransformationCompleted struct {
	_ byte // padding
}

// Experimental.
func NewTransformationCompleted() TransformationCompleted {
	_init_.Initialize()

	j := jsiiProxy_TransformationCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewTransformationCompleted_Override(t TransformationCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted",
		nil, // no parameters
		t,
	)
}

// EventBridge event pattern for Transformation Completed.
// Experimental.
func TransformationCompleted_TransformationCompletedPattern(options *TransformationCompleted_TransformationCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateTransformationCompleted_TransformationCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.events.TransformationCompleted",
		"transformationCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


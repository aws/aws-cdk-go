package previewawsb2bievents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.b2bi@AcknowledgementCompleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   acknowledgementCompleted := awscdkmixinspreview.Events.NewAcknowledgementCompleted()
//
// Experimental.
type AcknowledgementCompleted interface {
}

// The jsii proxy struct for AcknowledgementCompleted
type jsiiProxy_AcknowledgementCompleted struct {
	_ byte // padding
}

// Experimental.
func NewAcknowledgementCompleted() AcknowledgementCompleted {
	_init_.Initialize()

	j := jsiiProxy_AcknowledgementCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewAcknowledgementCompleted_Override(a AcknowledgementCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted",
		nil, // no parameters
		a,
	)
}

// EventBridge event pattern for Acknowledgement Completed.
// Experimental.
func AcknowledgementCompleted_AcknowledgementCompletedPattern(options *AcknowledgementCompleted_AcknowledgementCompletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateAcknowledgementCompleted_AcknowledgementCompletedPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_b2bi.events.AcknowledgementCompleted",
		"acknowledgementCompletedPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


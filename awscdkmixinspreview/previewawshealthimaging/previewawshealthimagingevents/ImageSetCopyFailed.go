package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetCopyFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopyFailed := awscdkmixinspreview.Events.NewImageSetCopyFailed()
//
// Experimental.
type ImageSetCopyFailed interface {
}

// The jsii proxy struct for ImageSetCopyFailed
type jsiiProxy_ImageSetCopyFailed struct {
	_ byte // padding
}

// Experimental.
func NewImageSetCopyFailed() ImageSetCopyFailed {
	_init_.Initialize()

	j := jsiiProxy_ImageSetCopyFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetCopyFailed_Override(i ImageSetCopyFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyFailed",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Copy Failed.
// Experimental.
func ImageSetCopyFailed_EventPattern(options *ImageSetCopyFailed_ImageSetCopyFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetCopyFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopyFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


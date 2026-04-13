package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetUpdateFailed.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetUpdateFailed := awscdkmixinspreview.Events.NewImageSetUpdateFailed()
//
// Experimental.
type ImageSetUpdateFailed interface {
}

// The jsii proxy struct for ImageSetUpdateFailed
type jsiiProxy_ImageSetUpdateFailed struct {
	_ byte // padding
}

// Experimental.
func NewImageSetUpdateFailed() ImageSetUpdateFailed {
	_init_.Initialize()

	j := jsiiProxy_ImageSetUpdateFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdateFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetUpdateFailed_Override(i ImageSetUpdateFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdateFailed",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Update Failed.
// Experimental.
func ImageSetUpdateFailed_EventPattern(options *ImageSetUpdateFailed_ImageSetUpdateFailedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetUpdateFailed_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdateFailed",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


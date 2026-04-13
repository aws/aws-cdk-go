package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetCopied.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopied := awscdkmixinspreview.Events.NewImageSetCopied()
//
// Experimental.
type ImageSetCopied interface {
}

// The jsii proxy struct for ImageSetCopied
type jsiiProxy_ImageSetCopied struct {
	_ byte // padding
}

// Experimental.
func NewImageSetCopied() ImageSetCopied {
	_init_.Initialize()

	j := jsiiProxy_ImageSetCopied{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopied",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetCopied_Override(i ImageSetCopied) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopied",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Copied.
// Experimental.
func ImageSetCopied_EventPattern(options *ImageSetCopied_ImageSetCopiedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetCopied_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopied",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


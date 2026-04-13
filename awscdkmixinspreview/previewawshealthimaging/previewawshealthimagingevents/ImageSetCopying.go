package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetCopying.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopying := awscdkmixinspreview.Events.NewImageSetCopying()
//
// Experimental.
type ImageSetCopying interface {
}

// The jsii proxy struct for ImageSetCopying
type jsiiProxy_ImageSetCopying struct {
	_ byte // padding
}

// Experimental.
func NewImageSetCopying() ImageSetCopying {
	_init_.Initialize()

	j := jsiiProxy_ImageSetCopying{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopying",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetCopying_Override(i ImageSetCopying) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopying",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Copying.
// Experimental.
func ImageSetCopying_EventPattern(options *ImageSetCopying_ImageSetCopyingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetCopying_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCopying",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


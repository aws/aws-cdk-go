package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetUpdating.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetUpdating := awscdkmixinspreview.Events.NewImageSetUpdating()
//
// Experimental.
type ImageSetUpdating interface {
}

// The jsii proxy struct for ImageSetUpdating
type jsiiProxy_ImageSetUpdating struct {
	_ byte // padding
}

// Experimental.
func NewImageSetUpdating() ImageSetUpdating {
	_init_.Initialize()

	j := jsiiProxy_ImageSetUpdating{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdating",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetUpdating_Override(i ImageSetUpdating) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdating",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Updating.
// Experimental.
func ImageSetUpdating_ImageSetUpdatingPattern(options *ImageSetUpdating_ImageSetUpdatingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetUpdating_ImageSetUpdatingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdating",
		"imageSetUpdatingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetUpdated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetUpdated := awscdkmixinspreview.Events.NewImageSetUpdated()
//
// Experimental.
type ImageSetUpdated interface {
}

// The jsii proxy struct for ImageSetUpdated
type jsiiProxy_ImageSetUpdated struct {
	_ byte // padding
}

// Experimental.
func NewImageSetUpdated() ImageSetUpdated {
	_init_.Initialize()

	j := jsiiProxy_ImageSetUpdated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetUpdated_Override(i ImageSetUpdated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdated",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Updated.
// Experimental.
func ImageSetUpdated_EventPattern(options *ImageSetUpdated_ImageSetUpdatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetUpdated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetUpdated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


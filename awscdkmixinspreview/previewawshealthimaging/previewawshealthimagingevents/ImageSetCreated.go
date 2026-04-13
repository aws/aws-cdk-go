package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetCreated.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCreated := awscdkmixinspreview.Events.NewImageSetCreated()
//
// Experimental.
type ImageSetCreated interface {
}

// The jsii proxy struct for ImageSetCreated
type jsiiProxy_ImageSetCreated struct {
	_ byte // padding
}

// Experimental.
func NewImageSetCreated() ImageSetCreated {
	_init_.Initialize()

	j := jsiiProxy_ImageSetCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetCreated_Override(i ImageSetCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCreated",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Created.
// Experimental.
func ImageSetCreated_EventPattern(options *ImageSetCreated_ImageSetCreatedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetCreated_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetCreated",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


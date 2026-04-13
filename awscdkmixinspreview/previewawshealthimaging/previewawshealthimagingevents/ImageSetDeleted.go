package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetDeleted.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetDeleted := awscdkmixinspreview.Events.NewImageSetDeleted()
//
// Experimental.
type ImageSetDeleted interface {
}

// The jsii proxy struct for ImageSetDeleted
type jsiiProxy_ImageSetDeleted struct {
	_ byte // padding
}

// Experimental.
func NewImageSetDeleted() ImageSetDeleted {
	_init_.Initialize()

	j := jsiiProxy_ImageSetDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetDeleted_Override(i ImageSetDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleted",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Deleted.
// Experimental.
func ImageSetDeleted_EventPattern(options *ImageSetDeleted_ImageSetDeletedProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetDeleted_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleted",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


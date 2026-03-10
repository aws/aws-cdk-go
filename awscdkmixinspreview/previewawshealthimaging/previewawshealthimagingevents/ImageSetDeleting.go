package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.healthimaging@ImageSetDeleting.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetDeleting := awscdkmixinspreview.Events.NewImageSetDeleting()
//
// Experimental.
type ImageSetDeleting interface {
}

// The jsii proxy struct for ImageSetDeleting
type jsiiProxy_ImageSetDeleting struct {
	_ byte // padding
}

// Experimental.
func NewImageSetDeleting() ImageSetDeleting {
	_init_.Initialize()

	j := jsiiProxy_ImageSetDeleting{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewImageSetDeleting_Override(i ImageSetDeleting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleting",
		nil, // no parameters
		i,
	)
}

// EventBridge event pattern for Image Set Deleting.
// Experimental.
func ImageSetDeleting_ImageSetDeletingPattern(options *ImageSetDeleting_ImageSetDeletingProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateImageSetDeleting_ImageSetDeletingPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.ImageSetDeleting",
		"imageSetDeletingPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


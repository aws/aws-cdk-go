package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetUpdating event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetUpdating := #error#.NewImageSetUpdating()
//
// Experimental.
type DatastoreEvents_ImageSetUpdating interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetUpdating
type jsiiProxy_DatastoreEvents_ImageSetUpdating struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetUpdating() DatastoreEvents_ImageSetUpdating {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetUpdating{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdating",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetUpdating_Override(d DatastoreEvents_ImageSetUpdating) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdating",
		nil, // no parameters
		d,
	)
}


package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetUpdated event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetUpdated := #error#.NewImageSetUpdated()
//
// Experimental.
type DatastoreEvents_ImageSetUpdated interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetUpdated
type jsiiProxy_DatastoreEvents_ImageSetUpdated struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetUpdated() DatastoreEvents_ImageSetUpdated {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetUpdated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetUpdated_Override(d DatastoreEvents_ImageSetUpdated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdated",
		nil, // no parameters
		d,
	)
}


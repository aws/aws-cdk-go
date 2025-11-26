package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetDeleting event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetDeleting := #error#.NewImageSetDeleting()
//
// Experimental.
type DatastoreEvents_ImageSetDeleting interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetDeleting
type jsiiProxy_DatastoreEvents_ImageSetDeleting struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetDeleting() DatastoreEvents_ImageSetDeleting {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetDeleting{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetDeleting_Override(d DatastoreEvents_ImageSetDeleting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleting",
		nil, // no parameters
		d,
	)
}


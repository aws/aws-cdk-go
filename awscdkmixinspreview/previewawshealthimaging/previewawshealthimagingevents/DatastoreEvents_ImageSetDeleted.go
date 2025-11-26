package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetDeleted event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetDeleted := #error#.NewImageSetDeleted()
//
// Experimental.
type DatastoreEvents_ImageSetDeleted interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetDeleted
type jsiiProxy_DatastoreEvents_ImageSetDeleted struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetDeleted() DatastoreEvents_ImageSetDeleted {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetDeleted_Override(d DatastoreEvents_ImageSetDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetDeleted",
		nil, // no parameters
		d,
	)
}


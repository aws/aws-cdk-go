package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetCreated event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCreated := #error#.NewImageSetCreated()
//
// Experimental.
type DatastoreEvents_ImageSetCreated interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetCreated
type jsiiProxy_DatastoreEvents_ImageSetCreated struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetCreated() DatastoreEvents_ImageSetCreated {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetCreated_Override(d DatastoreEvents_ImageSetCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCreated",
		nil, // no parameters
		d,
	)
}


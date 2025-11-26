package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetCopying event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopying := #error#.NewImageSetCopying()
//
// Experimental.
type DatastoreEvents_ImageSetCopying interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetCopying
type jsiiProxy_DatastoreEvents_ImageSetCopying struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetCopying() DatastoreEvents_ImageSetCopying {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetCopying{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopying",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetCopying_Override(d DatastoreEvents_ImageSetCopying) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopying",
		nil, // no parameters
		d,
	)
}


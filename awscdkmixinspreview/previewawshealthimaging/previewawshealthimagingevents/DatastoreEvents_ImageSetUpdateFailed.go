package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetUpdateFailed event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetUpdateFailed := #error#.NewImageSetUpdateFailed()
//
// Experimental.
type DatastoreEvents_ImageSetUpdateFailed interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetUpdateFailed
type jsiiProxy_DatastoreEvents_ImageSetUpdateFailed struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetUpdateFailed() DatastoreEvents_ImageSetUpdateFailed {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetUpdateFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdateFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetUpdateFailed_Override(d DatastoreEvents_ImageSetUpdateFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetUpdateFailed",
		nil, // no parameters
		d,
	)
}


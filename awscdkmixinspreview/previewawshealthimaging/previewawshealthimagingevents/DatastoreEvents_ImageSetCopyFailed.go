package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetCopyFailed event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopyFailed := #error#.NewImageSetCopyFailed()
//
// Experimental.
type DatastoreEvents_ImageSetCopyFailed interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetCopyFailed
type jsiiProxy_DatastoreEvents_ImageSetCopyFailed struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetCopyFailed() DatastoreEvents_ImageSetCopyFailed {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetCopyFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetCopyFailed_Override(d DatastoreEvents_ImageSetCopyFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyFailed",
		nil, // no parameters
		d,
	)
}


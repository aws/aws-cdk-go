package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImageSetCopyingWithReadOnlyAccess event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   imageSetCopyingWithReadOnlyAccess := #error#.NewImageSetCopyingWithReadOnlyAccess()
//
// Experimental.
type DatastoreEvents_ImageSetCopyingWithReadOnlyAccess interface {
}

// The jsii proxy struct for DatastoreEvents_ImageSetCopyingWithReadOnlyAccess
type jsiiProxy_DatastoreEvents_ImageSetCopyingWithReadOnlyAccess struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImageSetCopyingWithReadOnlyAccess() DatastoreEvents_ImageSetCopyingWithReadOnlyAccess {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImageSetCopyingWithReadOnlyAccess{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyingWithReadOnlyAccess",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImageSetCopyingWithReadOnlyAccess_Override(d DatastoreEvents_ImageSetCopyingWithReadOnlyAccess) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImageSetCopyingWithReadOnlyAccess",
		nil, // no parameters
		d,
	)
}


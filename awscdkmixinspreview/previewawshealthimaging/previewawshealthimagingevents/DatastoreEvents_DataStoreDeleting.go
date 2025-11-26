package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@DataStoreDeleting event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreDeleting := #error#.NewDataStoreDeleting()
//
// Experimental.
type DatastoreEvents_DataStoreDeleting interface {
}

// The jsii proxy struct for DatastoreEvents_DataStoreDeleting
type jsiiProxy_DatastoreEvents_DataStoreDeleting struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_DataStoreDeleting() DatastoreEvents_DataStoreDeleting {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_DataStoreDeleting{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleting",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_DataStoreDeleting_Override(d DatastoreEvents_DataStoreDeleting) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleting",
		nil, // no parameters
		d,
	)
}


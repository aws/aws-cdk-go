package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@DataStoreCreating event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreating := #error#.NewDataStoreCreating()
//
// Experimental.
type DatastoreEvents_DataStoreCreating interface {
}

// The jsii proxy struct for DatastoreEvents_DataStoreCreating
type jsiiProxy_DatastoreEvents_DataStoreCreating struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_DataStoreCreating() DatastoreEvents_DataStoreCreating {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_DataStoreCreating{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreating",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_DataStoreCreating_Override(d DatastoreEvents_DataStoreCreating) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreating",
		nil, // no parameters
		d,
	)
}


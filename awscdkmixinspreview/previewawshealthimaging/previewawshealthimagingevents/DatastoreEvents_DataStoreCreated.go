package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@DataStoreCreated event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreated := #error#.NewDataStoreCreated()
//
// Experimental.
type DatastoreEvents_DataStoreCreated interface {
}

// The jsii proxy struct for DatastoreEvents_DataStoreCreated
type jsiiProxy_DatastoreEvents_DataStoreCreated struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_DataStoreCreated() DatastoreEvents_DataStoreCreated {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_DataStoreCreated{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreated",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_DataStoreCreated_Override(d DatastoreEvents_DataStoreCreated) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreated",
		nil, // no parameters
		d,
	)
}


package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@DataStoreCreationFailed event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreCreationFailed := #error#.NewDataStoreCreationFailed()
//
// Experimental.
type DatastoreEvents_DataStoreCreationFailed interface {
}

// The jsii proxy struct for DatastoreEvents_DataStoreCreationFailed
type jsiiProxy_DatastoreEvents_DataStoreCreationFailed struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_DataStoreCreationFailed() DatastoreEvents_DataStoreCreationFailed {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_DataStoreCreationFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreationFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_DataStoreCreationFailed_Override(d DatastoreEvents_DataStoreCreationFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreCreationFailed",
		nil, // no parameters
		d,
	)
}


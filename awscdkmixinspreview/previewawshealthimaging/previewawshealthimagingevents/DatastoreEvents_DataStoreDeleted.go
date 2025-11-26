package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@DataStoreDeleted event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   dataStoreDeleted := #error#.NewDataStoreDeleted()
//
// Experimental.
type DatastoreEvents_DataStoreDeleted interface {
}

// The jsii proxy struct for DatastoreEvents_DataStoreDeleted
type jsiiProxy_DatastoreEvents_DataStoreDeleted struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_DataStoreDeleted() DatastoreEvents_DataStoreDeleted {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_DataStoreDeleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_DataStoreDeleted_Override(d DatastoreEvents_DataStoreDeleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.DataStoreDeleted",
		nil, // no parameters
		d,
	)
}


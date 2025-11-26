package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImportJobInProgress event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobInProgress := #error#.NewImportJobInProgress()
//
// Experimental.
type DatastoreEvents_ImportJobInProgress interface {
}

// The jsii proxy struct for DatastoreEvents_ImportJobInProgress
type jsiiProxy_DatastoreEvents_ImportJobInProgress struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImportJobInProgress() DatastoreEvents_ImportJobInProgress {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImportJobInProgress{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobInProgress",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImportJobInProgress_Override(d DatastoreEvents_ImportJobInProgress) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobInProgress",
		nil, // no parameters
		d,
	)
}


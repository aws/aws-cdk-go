package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImportJobCompleted event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobCompleted := #error#.NewImportJobCompleted()
//
// Experimental.
type DatastoreEvents_ImportJobCompleted interface {
}

// The jsii proxy struct for DatastoreEvents_ImportJobCompleted
type jsiiProxy_DatastoreEvents_ImportJobCompleted struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImportJobCompleted() DatastoreEvents_ImportJobCompleted {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImportJobCompleted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobCompleted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImportJobCompleted_Override(d DatastoreEvents_ImportJobCompleted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobCompleted",
		nil, // no parameters
		d,
	)
}


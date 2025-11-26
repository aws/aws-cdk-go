package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImportJobFailed event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobFailed := #error#.NewImportJobFailed()
//
// Experimental.
type DatastoreEvents_ImportJobFailed interface {
}

// The jsii proxy struct for DatastoreEvents_ImportJobFailed
type jsiiProxy_DatastoreEvents_ImportJobFailed struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImportJobFailed() DatastoreEvents_ImportJobFailed {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImportJobFailed{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobFailed",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImportJobFailed_Override(d DatastoreEvents_ImportJobFailed) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobFailed",
		nil, // no parameters
		d,
	)
}


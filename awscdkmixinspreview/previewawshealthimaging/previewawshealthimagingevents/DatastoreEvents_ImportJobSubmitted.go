package previewawshealthimagingevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.healthimaging@ImportJobSubmitted event types for Datastore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   importJobSubmitted := #error#.NewImportJobSubmitted()
//
// Experimental.
type DatastoreEvents_ImportJobSubmitted interface {
}

// The jsii proxy struct for DatastoreEvents_ImportJobSubmitted
type jsiiProxy_DatastoreEvents_ImportJobSubmitted struct {
	_ byte // padding
}

// Experimental.
func NewDatastoreEvents_ImportJobSubmitted() DatastoreEvents_ImportJobSubmitted {
	_init_.Initialize()

	j := jsiiProxy_DatastoreEvents_ImportJobSubmitted{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobSubmitted",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewDatastoreEvents_ImportJobSubmitted_Override(d DatastoreEvents_ImportJobSubmitted) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_healthimaging.events.DatastoreEvents.ImportJobSubmitted",
		nil, // no parameters
		d,
	)
}


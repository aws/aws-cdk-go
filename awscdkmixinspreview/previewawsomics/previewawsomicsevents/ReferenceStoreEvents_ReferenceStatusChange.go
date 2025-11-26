package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.omics@ReferenceStatusChange event types for ReferenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceStatusChange := #error#.NewReferenceStatusChange()
//
// Experimental.
type ReferenceStoreEvents_ReferenceStatusChange interface {
}

// The jsii proxy struct for ReferenceStoreEvents_ReferenceStatusChange
type jsiiProxy_ReferenceStoreEvents_ReferenceStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReferenceStoreEvents_ReferenceStatusChange() ReferenceStoreEvents_ReferenceStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReferenceStoreEvents_ReferenceStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReferenceStoreEvents_ReferenceStatusChange_Override(r ReferenceStoreEvents_ReferenceStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceStatusChange",
		nil, // no parameters
		r,
	)
}


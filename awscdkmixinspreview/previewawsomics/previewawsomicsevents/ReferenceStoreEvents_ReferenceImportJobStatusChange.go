package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.omics@ReferenceImportJobStatusChange event types for ReferenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceImportJobStatusChange := #error#.NewReferenceImportJobStatusChange()
//
// Experimental.
type ReferenceStoreEvents_ReferenceImportJobStatusChange interface {
}

// The jsii proxy struct for ReferenceStoreEvents_ReferenceImportJobStatusChange
type jsiiProxy_ReferenceStoreEvents_ReferenceImportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReferenceStoreEvents_ReferenceImportJobStatusChange() ReferenceStoreEvents_ReferenceImportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReferenceStoreEvents_ReferenceImportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceImportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReferenceStoreEvents_ReferenceImportJobStatusChange_Override(r ReferenceStoreEvents_ReferenceImportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStoreEvents.ReferenceImportJobStatusChange",
		nil, // no parameters
		r,
	)
}


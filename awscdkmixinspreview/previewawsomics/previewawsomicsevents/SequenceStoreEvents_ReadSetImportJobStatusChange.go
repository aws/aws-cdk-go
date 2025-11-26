package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.omics@ReadSetImportJobStatusChange event types for SequenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetImportJobStatusChange := #error#.NewReadSetImportJobStatusChange()
//
// Experimental.
type SequenceStoreEvents_ReadSetImportJobStatusChange interface {
}

// The jsii proxy struct for SequenceStoreEvents_ReadSetImportJobStatusChange
type jsiiProxy_SequenceStoreEvents_ReadSetImportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSequenceStoreEvents_ReadSetImportJobStatusChange() SequenceStoreEvents_ReadSetImportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SequenceStoreEvents_ReadSetImportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetImportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSequenceStoreEvents_ReadSetImportJobStatusChange_Override(s SequenceStoreEvents_ReadSetImportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetImportJobStatusChange",
		nil, // no parameters
		s,
	)
}


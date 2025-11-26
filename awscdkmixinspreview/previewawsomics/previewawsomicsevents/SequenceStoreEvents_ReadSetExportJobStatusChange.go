package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.omics@ReadSetExportJobStatusChange event types for SequenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetExportJobStatusChange := #error#.NewReadSetExportJobStatusChange()
//
// Experimental.
type SequenceStoreEvents_ReadSetExportJobStatusChange interface {
}

// The jsii proxy struct for SequenceStoreEvents_ReadSetExportJobStatusChange
type jsiiProxy_SequenceStoreEvents_ReadSetExportJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSequenceStoreEvents_ReadSetExportJobStatusChange() SequenceStoreEvents_ReadSetExportJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SequenceStoreEvents_ReadSetExportJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetExportJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSequenceStoreEvents_ReadSetExportJobStatusChange_Override(s SequenceStoreEvents_ReadSetExportJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetExportJobStatusChange",
		nil, // no parameters
		s,
	)
}


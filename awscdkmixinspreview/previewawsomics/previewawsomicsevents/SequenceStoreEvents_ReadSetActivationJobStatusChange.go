package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.omics@ReadSetActivationJobStatusChange event types for SequenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetActivationJobStatusChange := #error#.NewReadSetActivationJobStatusChange()
//
// Experimental.
type SequenceStoreEvents_ReadSetActivationJobStatusChange interface {
}

// The jsii proxy struct for SequenceStoreEvents_ReadSetActivationJobStatusChange
type jsiiProxy_SequenceStoreEvents_ReadSetActivationJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSequenceStoreEvents_ReadSetActivationJobStatusChange() SequenceStoreEvents_ReadSetActivationJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SequenceStoreEvents_ReadSetActivationJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetActivationJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSequenceStoreEvents_ReadSetActivationJobStatusChange_Override(s SequenceStoreEvents_ReadSetActivationJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetActivationJobStatusChange",
		nil, // no parameters
		s,
	)
}


package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// aws.omics@ReadSetStatusChange event types for SequenceStore.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetStatusChange := #error#.NewReadSetStatusChange()
//
// Experimental.
type SequenceStoreEvents_ReadSetStatusChange interface {
}

// The jsii proxy struct for SequenceStoreEvents_ReadSetStatusChange
type jsiiProxy_SequenceStoreEvents_ReadSetStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSequenceStoreEvents_ReadSetStatusChange() SequenceStoreEvents_ReadSetStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SequenceStoreEvents_ReadSetStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSequenceStoreEvents_ReadSetStatusChange_Override(s SequenceStoreEvents_ReadSetStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreEvents.ReadSetStatusChange",
		nil, // no parameters
		s,
	)
}


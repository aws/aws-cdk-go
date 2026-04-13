package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@SequenceStoreStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   sequenceStoreStatusChange := awscdkmixinspreview.Events.NewSequenceStoreStatusChange()
//
// Experimental.
type SequenceStoreStatusChange interface {
}

// The jsii proxy struct for SequenceStoreStatusChange
type jsiiProxy_SequenceStoreStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSequenceStoreStatusChange() SequenceStoreStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SequenceStoreStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSequenceStoreStatusChange_Override(s SequenceStoreStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreStatusChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Sequence Store Status Change.
// Experimental.
func SequenceStoreStatusChange_EventPattern(options *SequenceStoreStatusChange_SequenceStoreStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSequenceStoreStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.SequenceStoreStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


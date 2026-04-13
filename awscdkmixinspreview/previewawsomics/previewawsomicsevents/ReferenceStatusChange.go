package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReferenceStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   referenceStatusChange := awscdkmixinspreview.Events.NewReferenceStatusChange()
//
// Experimental.
type ReferenceStatusChange interface {
}

// The jsii proxy struct for ReferenceStatusChange
type jsiiProxy_ReferenceStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReferenceStatusChange() ReferenceStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReferenceStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReferenceStatusChange_Override(r ReferenceStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Reference Status Change.
// Experimental.
func ReferenceStatusChange_EventPattern(options *ReferenceStatusChange_ReferenceStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReferenceStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReferenceStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


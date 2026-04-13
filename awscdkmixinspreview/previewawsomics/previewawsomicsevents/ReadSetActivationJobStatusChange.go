package previewawsomicsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.omics@ReadSetActivationJobStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   readSetActivationJobStatusChange := awscdkmixinspreview.Events.NewReadSetActivationJobStatusChange()
//
// Experimental.
type ReadSetActivationJobStatusChange interface {
}

// The jsii proxy struct for ReadSetActivationJobStatusChange
type jsiiProxy_ReadSetActivationJobStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewReadSetActivationJobStatusChange() ReadSetActivationJobStatusChange {
	_init_.Initialize()

	j := jsiiProxy_ReadSetActivationJobStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetActivationJobStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewReadSetActivationJobStatusChange_Override(r ReadSetActivationJobStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetActivationJobStatusChange",
		nil, // no parameters
		r,
	)
}

// EventBridge event pattern for Read Set Activation Job Status Change.
// Experimental.
func ReadSetActivationJobStatusChange_EventPattern(options *ReadSetActivationJobStatusChange_ReadSetActivationJobStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateReadSetActivationJobStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_omics.events.ReadSetActivationJobStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


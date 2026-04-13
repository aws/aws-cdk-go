package previewawssyntheticsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.synthetics@SyntheticsCanaryStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   syntheticsCanaryStatusChange := awscdkmixinspreview.Events.NewSyntheticsCanaryStatusChange()
//
// Experimental.
type SyntheticsCanaryStatusChange interface {
}

// The jsii proxy struct for SyntheticsCanaryStatusChange
type jsiiProxy_SyntheticsCanaryStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewSyntheticsCanaryStatusChange() SyntheticsCanaryStatusChange {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSyntheticsCanaryStatusChange_Override(s SyntheticsCanaryStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Synthetics Canary Status Change.
// Experimental.
func SyntheticsCanaryStatusChange_EventPattern(options *SyntheticsCanaryStatusChange_SyntheticsCanaryStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSyntheticsCanaryStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


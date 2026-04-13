package previewawssyntheticsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.synthetics@SyntheticsCanaryTestRunFailure.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   syntheticsCanaryTestRunFailure := awscdkmixinspreview.Events.NewSyntheticsCanaryTestRunFailure()
//
// Experimental.
type SyntheticsCanaryTestRunFailure interface {
}

// The jsii proxy struct for SyntheticsCanaryTestRunFailure
type jsiiProxy_SyntheticsCanaryTestRunFailure struct {
	_ byte // padding
}

// Experimental.
func NewSyntheticsCanaryTestRunFailure() SyntheticsCanaryTestRunFailure {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryTestRunFailure{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunFailure",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSyntheticsCanaryTestRunFailure_Override(s SyntheticsCanaryTestRunFailure) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunFailure",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Synthetics Canary TestRun Failure.
// Experimental.
func SyntheticsCanaryTestRunFailure_EventPattern(options *SyntheticsCanaryTestRunFailure_SyntheticsCanaryTestRunFailureProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSyntheticsCanaryTestRunFailure_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunFailure",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


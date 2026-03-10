package previewawssyntheticsevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.synthetics@SyntheticsCanaryTestRunSuccessful.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   syntheticsCanaryTestRunSuccessful := awscdkmixinspreview.Events.NewSyntheticsCanaryTestRunSuccessful()
//
// Experimental.
type SyntheticsCanaryTestRunSuccessful interface {
}

// The jsii proxy struct for SyntheticsCanaryTestRunSuccessful
type jsiiProxy_SyntheticsCanaryTestRunSuccessful struct {
	_ byte // padding
}

// Experimental.
func NewSyntheticsCanaryTestRunSuccessful() SyntheticsCanaryTestRunSuccessful {
	_init_.Initialize()

	j := jsiiProxy_SyntheticsCanaryTestRunSuccessful{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunSuccessful",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewSyntheticsCanaryTestRunSuccessful_Override(s SyntheticsCanaryTestRunSuccessful) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunSuccessful",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Synthetics Canary TestRun Successful.
// Experimental.
func SyntheticsCanaryTestRunSuccessful_SyntheticsCanaryTestRunSuccessfulPattern(options *SyntheticsCanaryTestRunSuccessful_SyntheticsCanaryTestRunSuccessfulProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateSyntheticsCanaryTestRunSuccessful_SyntheticsCanaryTestRunSuccessfulPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_synthetics.events.SyntheticsCanaryTestRunSuccessful",
		"syntheticsCanaryTestRunSuccessfulPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


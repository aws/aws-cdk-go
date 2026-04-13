package previewawsfisevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.fis@FISExperimentStateChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   fISExperimentStateChange := awscdkmixinspreview.Events.NewFISExperimentStateChange()
//
// Experimental.
type FISExperimentStateChange interface {
}

// The jsii proxy struct for FISExperimentStateChange
type jsiiProxy_FISExperimentStateChange struct {
	_ byte // padding
}

// Experimental.
func NewFISExperimentStateChange() FISExperimentStateChange {
	_init_.Initialize()

	j := jsiiProxy_FISExperimentStateChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewFISExperimentStateChange_Override(f FISExperimentStateChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange",
		nil, // no parameters
		f,
	)
}

// EventBridge event pattern for FIS Experiment State Change.
// Experimental.
func FISExperimentStateChange_EventPattern(options *FISExperimentStateChange_FISExperimentStateChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateFISExperimentStateChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_fis.events.FISExperimentStateChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


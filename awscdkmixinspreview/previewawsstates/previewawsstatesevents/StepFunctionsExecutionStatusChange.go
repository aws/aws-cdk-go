package previewawsstatesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
)

// EventBridge event pattern for aws.states@StepFunctionsExecutionStatusChange.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//
//   stepFunctionsExecutionStatusChange := awscdkmixinspreview.Events.NewStepFunctionsExecutionStatusChange()
//
// Experimental.
type StepFunctionsExecutionStatusChange interface {
}

// The jsii proxy struct for StepFunctionsExecutionStatusChange
type jsiiProxy_StepFunctionsExecutionStatusChange struct {
	_ byte // padding
}

// Experimental.
func NewStepFunctionsExecutionStatusChange() StepFunctionsExecutionStatusChange {
	_init_.Initialize()

	j := jsiiProxy_StepFunctionsExecutionStatusChange{}

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_states.events.StepFunctionsExecutionStatusChange",
		nil, // no parameters
		&j,
	)

	return &j
}

// Experimental.
func NewStepFunctionsExecutionStatusChange_Override(s StepFunctionsExecutionStatusChange) {
	_init_.Initialize()

	_jsii_.Create(
		"@aws-cdk/mixins-preview.aws_states.events.StepFunctionsExecutionStatusChange",
		nil, // no parameters
		s,
	)
}

// EventBridge event pattern for Step Functions Execution Status Change.
// Experimental.
func StepFunctionsExecutionStatusChange_EventPattern(options *StepFunctionsExecutionStatusChange_StepFunctionsExecutionStatusChangeProps) *awsevents.EventPattern {
	_init_.Initialize()

	if err := validateStepFunctionsExecutionStatusChange_EventPatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_states.events.StepFunctionsExecutionStatusChange",
		"eventPattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


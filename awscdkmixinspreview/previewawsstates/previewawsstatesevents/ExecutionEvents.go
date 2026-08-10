package previewawsstatesevents

import (
	_init_ "github.com/aws/aws-cdk-go/awscdkmixinspreview/v2/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/interfaces/interfacesawsstates"
)

// EventBridge event patterns for Execution.
//
// Example:
//   // The code below shows an example of how to instantiate this type.
//   // The values are placeholders you should change.
//   import "github.com/aws/aws-cdk-go/awscdkmixinspreview"
//   import "github.com/aws/aws-cdk-go/awscdk"
//
//   var executionRef IExecutionRef
//
//   executionEvents := awscdkmixinspreview.Events.ExecutionEvents_FromExecution(executionRef)
//
// Experimental.
type ExecutionEvents interface {
	// EventBridge event pattern for Execution Step Functions Execution Status Change.
	// Experimental.
	StepFunctionsExecutionStatusChangePattern(options *StepFunctionsExecutionStatusChange_StepFunctionsExecutionStatusChangeProps) *awsevents.EventPattern
}

// The jsii proxy struct for ExecutionEvents
type jsiiProxy_ExecutionEvents struct {
	_ byte // padding
}

// Create ExecutionEvents from a Execution reference.
// Experimental.
func ExecutionEvents_FromExecution(executionRef interfacesawsstates.IExecutionRef) ExecutionEvents {
	_init_.Initialize()

	if err := validateExecutionEvents_FromExecutionParameters(executionRef); err != nil {
		panic(err)
	}
	var returns ExecutionEvents

	_jsii_.StaticInvoke(
		"@aws-cdk/mixins-preview.aws_states.events.ExecutionEvents",
		"fromExecution",
		[]interface{}{executionRef},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExecutionEvents) StepFunctionsExecutionStatusChangePattern(options *StepFunctionsExecutionStatusChange_StepFunctionsExecutionStatusChangeProps) *awsevents.EventPattern {
	if err := e.validateStepFunctionsExecutionStatusChangePatternParameters(options); err != nil {
		panic(err)
	}
	var returns *awsevents.EventPattern

	_jsii_.Invoke(
		e,
		"stepFunctionsExecutionStatusChangePattern",
		[]interface{}{options},
		&returns,
	)

	return returns
}


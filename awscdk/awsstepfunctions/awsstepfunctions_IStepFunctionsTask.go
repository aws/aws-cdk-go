package awsstepfunctions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Interface for resources that can be used as tasks.
// Deprecated: replaced by `TaskStateBase`.
type IStepFunctionsTask interface {
	// Called when the task object is used in a workflow.
	// Deprecated: replaced by `TaskStateBase`.
	Bind(task Task) *StepFunctionsTaskConfig
}

// The jsii proxy for IStepFunctionsTask
type jsiiProxy_IStepFunctionsTask struct {
	_ byte // padding
}

func (i *jsiiProxy_IStepFunctionsTask) Bind(task Task) *StepFunctionsTaskConfig {
	if err := i.validateBindParameters(task); err != nil {
		panic(err)
	}
	var returns *StepFunctionsTaskConfig

	_jsii_.Invoke(
		i,
		"bind",
		[]interface{}{task},
		&returns,
	)

	return returns
}


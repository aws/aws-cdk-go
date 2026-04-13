//go:build !no_runtime_type_checking

package previewawsimagebuilderevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEC2ImageBuilderWorkflowStepWaiting_EventPatternParameters(options *EC2ImageBuilderWorkflowStepWaiting_EC2ImageBuilderWorkflowStepWaitingProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


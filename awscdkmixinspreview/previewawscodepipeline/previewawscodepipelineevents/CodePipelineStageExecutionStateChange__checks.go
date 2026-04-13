//go:build !no_runtime_type_checking

package previewawscodepipelineevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCodePipelineStageExecutionStateChange_EventPatternParameters(options *CodePipelineStageExecutionStateChange_CodePipelineStageExecutionStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


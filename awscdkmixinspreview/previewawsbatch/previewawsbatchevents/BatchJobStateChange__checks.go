//go:build !no_runtime_type_checking

package previewawsbatchevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateBatchJobStateChange_EventPatternParameters(options *BatchJobStateChange_BatchJobStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


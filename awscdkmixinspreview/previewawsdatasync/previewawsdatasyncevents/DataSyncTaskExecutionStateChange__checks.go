//go:build !no_runtime_type_checking

package previewawsdatasyncevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDataSyncTaskExecutionStateChange_EventPatternParameters(options *DataSyncTaskExecutionStateChange_DataSyncTaskExecutionStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


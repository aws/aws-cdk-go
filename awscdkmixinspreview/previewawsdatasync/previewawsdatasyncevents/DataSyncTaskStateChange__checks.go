//go:build !no_runtime_type_checking

package previewawsdatasyncevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDataSyncTaskStateChange_EventPatternParameters(options *DataSyncTaskStateChange_DataSyncTaskStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


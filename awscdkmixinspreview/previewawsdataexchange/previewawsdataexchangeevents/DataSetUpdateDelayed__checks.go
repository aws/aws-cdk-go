//go:build !no_runtime_type_checking

package previewawsdataexchangeevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDataSetUpdateDelayed_EventPatternParameters(options *DataSetUpdateDelayed_DataSetUpdateDelayedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


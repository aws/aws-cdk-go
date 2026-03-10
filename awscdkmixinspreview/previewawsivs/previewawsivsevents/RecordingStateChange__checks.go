//go:build !no_runtime_type_checking

package previewawsivsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateRecordingStateChange_RecordingStateChangePatternParameters(options *RecordingStateChange_RecordingStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


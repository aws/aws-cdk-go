//go:build !no_runtime_type_checking

package previewawshealthlakeevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateImportJobCompleted_ImportJobCompletedPatternParameters(options *ImportJobCompleted_ImportJobCompletedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


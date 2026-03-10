//go:build !no_runtime_type_checking

package previewawshealthimagingevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateImageSetUpdateFailed_ImageSetUpdateFailedPatternParameters(options *ImageSetUpdateFailed_ImageSetUpdateFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


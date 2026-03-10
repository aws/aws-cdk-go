//go:build !no_runtime_type_checking

package previewawsb2bievents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateTransformationFailed_TransformationFailedPatternParameters(options *TransformationFailed_TransformationFailedProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

package previewawssyntheticsevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateSyntheticsCanaryTestRunSuccessful_SyntheticsCanaryTestRunSuccessfulPatternParameters(options *SyntheticsCanaryTestRunSuccessful_SyntheticsCanaryTestRunSuccessfulProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


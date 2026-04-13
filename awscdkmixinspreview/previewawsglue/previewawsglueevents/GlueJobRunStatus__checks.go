//go:build !no_runtime_type_checking

package previewawsglueevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateGlueJobRunStatus_EventPatternParameters(options *GlueJobRunStatus_GlueJobRunStatusProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

package previewawstranscribeevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCallAnalyticsJobStateChange_EventPatternParameters(options *CallAnalyticsJobStateChange_CallAnalyticsJobStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


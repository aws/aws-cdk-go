//go:build !no_runtime_type_checking

package previewawsssmevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateCalendarStateChange_CalendarStateChangePatternParameters(options *CalendarStateChange_CalendarStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


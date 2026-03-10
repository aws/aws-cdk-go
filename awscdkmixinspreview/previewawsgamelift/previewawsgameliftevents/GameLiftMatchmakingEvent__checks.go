//go:build !no_runtime_type_checking

package previewawsgameliftevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateGameLiftMatchmakingEvent_GameLiftMatchmakingEventPatternParameters(options *GameLiftMatchmakingEvent_GameLiftMatchmakingEventProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


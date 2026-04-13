//go:build !no_runtime_type_checking

package previewawsmedialiveevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMediaLiveMultiplexStateChange_EventPatternParameters(options *MediaLiveMultiplexStateChange_MediaLiveMultiplexStateChangeProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


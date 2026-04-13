//go:build !no_runtime_type_checking

package previewawsmedialiveevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateMediaLiveChannelAlert_EventPatternParameters(options *MediaLiveChannelAlert_MediaLiveChannelAlertProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


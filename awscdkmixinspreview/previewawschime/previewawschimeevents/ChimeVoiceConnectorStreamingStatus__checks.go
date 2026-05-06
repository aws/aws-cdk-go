//go:build !no_runtime_type_checking

package previewawschimeevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateChimeVoiceConnectorStreamingStatus_EventPatternParameters(options *ChimeVoiceConnectorStreamingStatus_ChimeVoiceConnectorStreamingStatusProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


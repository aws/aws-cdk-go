//go:build !no_runtime_type_checking

package previewawsvoiceidevents

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateVoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionPatternParameters(options *VoiceIdSessionSpeakerEnrollmentAction_VoiceIdSessionSpeakerEnrollmentActionProps) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateHlsSettings_AudioOnlyParameters(props *AudioOnlyHlsSettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateHlsSettings_Fmp4Parameters(props *Fmp4HlsSettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateHlsSettings_StandardParameters(props *StandardHlsSettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


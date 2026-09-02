//go:build !no_runtime_type_checking

package awsmedialivealpha

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateAudioCodecSettings_AacParameters(props *AacSettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAudioCodecSettings_Ac3Parameters(props *Ac3SettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAudioCodecSettings_Eac3Parameters(props *Eac3SettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAudioCodecSettings_Eac3AtmosParameters(props *Eac3AtmosSettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAudioCodecSettings_Mp2Parameters(props *Mp2SettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateAudioCodecSettings_WavParameters(props *WavSettingsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}


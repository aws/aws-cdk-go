//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateAudioCodecSettings_AacParameters(props *AacSettingsProps) error {
	return nil
}

func validateAudioCodecSettings_Ac3Parameters(props *Ac3SettingsProps) error {
	return nil
}

func validateAudioCodecSettings_Eac3Parameters(props *Eac3SettingsProps) error {
	return nil
}

func validateAudioCodecSettings_Eac3AtmosParameters(props *Eac3AtmosSettingsProps) error {
	return nil
}

func validateAudioCodecSettings_Mp2Parameters(props *Mp2SettingsProps) error {
	return nil
}

func validateAudioCodecSettings_WavParameters(props *WavSettingsProps) error {
	return nil
}


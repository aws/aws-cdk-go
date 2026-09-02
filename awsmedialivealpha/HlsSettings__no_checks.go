//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateHlsSettings_AudioOnlyParameters(props *AudioOnlyHlsSettingsProps) error {
	return nil
}

func validateHlsSettings_Fmp4Parameters(props *Fmp4HlsSettingsProps) error {
	return nil
}

func validateHlsSettings_StandardParameters(props *StandardHlsSettingsProps) error {
	return nil
}


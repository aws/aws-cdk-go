//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateAudioSelector_ByLanguageParameters(name *string, languageCode *string) error {
	return nil
}

func validateAudioSelector_ByPidParameters(name *string, pids *[]*AudioPidConfig) error {
	return nil
}

func validateAudioSelector_ByTrackParameters(name *string, tracks *[]*AudioTrackConfig) error {
	return nil
}

func validateAudioSelector_DefaultParameters(name *string) error {
	return nil
}

func validateAudioSelector_HlsRenditionParameters(name *string, options *HlsRenditionSelectionOptions) error {
	return nil
}

func validateNewAudioSelectorParameters(name *string) error {
	return nil
}


//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateCaptionSelector_AncillaryParameters(name *string, options *AncillaryCaptionSourceOptions) error {
	return nil
}

func validateCaptionSelector_AribParameters(name *string) error {
	return nil
}

func validateCaptionSelector_ByLanguageParameters(name *string, languageCode *string) error {
	return nil
}

func validateCaptionSelector_DvbSubParameters(name *string, options *DvbSubCaptionSourceOptions) error {
	return nil
}

func validateCaptionSelector_EmbeddedParameters(name *string, options *EmbeddedCaptionSourceOptions) error {
	return nil
}

func validateCaptionSelector_Scte20Parameters(name *string, options *Scte20CaptionSourceOptions) error {
	return nil
}

func validateCaptionSelector_Scte27Parameters(name *string, options *Scte27CaptionSourceOptions) error {
	return nil
}

func validateCaptionSelector_SmartSubtitleParameters(name *string, options *SmartSubtitleSourceOptions) error {
	return nil
}

func validateCaptionSelector_TeletextParameters(name *string, options *TeletextCaptionSourceOptions) error {
	return nil
}

func validateNewCaptionSelectorParameters(name *string) error {
	return nil
}


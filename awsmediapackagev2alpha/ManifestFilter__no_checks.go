//go:build no_runtime_type_checking

package awsmediapackagev2alpha

// Building without runtime type checking enabled, so all the below just return nil

func validateManifestFilter_AudioCodecParameters(value AudioCodec) error {
	return nil
}

func validateManifestFilter_AudioCodecListParameters(values *[]AudioCodec) error {
	return nil
}

func validateManifestFilter_BitrateParameters(key BitrateFilterKey, value awscdk.Bitrate) error {
	return nil
}

func validateManifestFilter_BitrateComboParameters(key BitrateFilterKey, expressions *[]BitrateExpression) error {
	return nil
}

func validateManifestFilter_BitrateRangeParameters(key BitrateFilterKey, start awscdk.Bitrate, end awscdk.Bitrate) error {
	return nil
}

func validateManifestFilter_CustomParameters(custom *string) error {
	return nil
}

func validateManifestFilter_NumericParameters(key NumericFilterKey, value *float64) error {
	return nil
}

func validateManifestFilter_NumericComboParameters(key NumericFilterKey, expressions *[]NumericExpression) error {
	return nil
}

func validateManifestFilter_NumericListParameters(key NumericFilterKey, values *[]*float64) error {
	return nil
}

func validateManifestFilter_NumericRangeParameters(key NumericFilterKey, start *float64, end *float64) error {
	return nil
}

func validateManifestFilter_TextParameters(key TextFilterKey, value *string) error {
	return nil
}

func validateManifestFilter_TextListParameters(key TextFilterKey, values *[]*string) error {
	return nil
}

func validateManifestFilter_TrickplayTypeParameters(value TrickplayType) error {
	return nil
}

func validateManifestFilter_TrickplayTypeListParameters(values *[]TrickplayType) error {
	return nil
}

func validateManifestFilter_VideoCodecParameters(value VideoCodec) error {
	return nil
}

func validateManifestFilter_VideoCodecListParameters(values *[]VideoCodec) error {
	return nil
}

func validateManifestFilter_VideoDynamicRangeParameters(value VideoDynamicRange) error {
	return nil
}

func validateManifestFilter_VideoDynamicRangeListParameters(values *[]VideoDynamicRange) error {
	return nil
}

func validateNewManifestFilterParameters(filterString *string) error {
	return nil
}


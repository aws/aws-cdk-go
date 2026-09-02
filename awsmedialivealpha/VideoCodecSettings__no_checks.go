//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateVideoCodecSettings_Av1Parameters(props *Av1SettingsProps) error {
	return nil
}

func validateVideoCodecSettings_FrameCaptureParameters(props *FrameCaptureSettingsProps) error {
	return nil
}

func validateVideoCodecSettings_H264Parameters(props *H264SettingsProps) error {
	return nil
}

func validateVideoCodecSettings_H265Parameters(props *H265SettingsProps) error {
	return nil
}


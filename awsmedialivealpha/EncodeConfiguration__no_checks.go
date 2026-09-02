//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateEncodeConfiguration_AudioParameters(props *AudioEncodeProps) error {
	return nil
}

func validateEncodeConfiguration_CaptionParameters(props *CaptionEncodeProps) error {
	return nil
}

func validateEncodeConfiguration_VideoParameters(props *VideoEncodeProps) error {
	return nil
}


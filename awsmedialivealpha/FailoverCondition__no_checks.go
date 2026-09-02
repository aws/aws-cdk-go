//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateFailoverCondition_AudioSilenceParameters(props *AudioSilenceFailoverProps) error {
	return nil
}

func validateFailoverCondition_InputLossParameters(props *InputLossFailoverProps) error {
	return nil
}

func validateFailoverCondition_VideoBlackParameters(props *VideoBlackFailoverProps) error {
	return nil
}


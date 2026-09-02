//go:build no_runtime_type_checking

package awsmedialivealpha

// Building without runtime type checking enabled, so all the below just return nil

func validateH264RateControl_CbrParameters(props *CbrRateControlProps) error {
	return nil
}

func validateH264RateControl_QvbrParameters(props *QvbrRateControlProps) error {
	return nil
}

func validateH264RateControl_VbrParameters(props *VbrRateControlProps) error {
	return nil
}


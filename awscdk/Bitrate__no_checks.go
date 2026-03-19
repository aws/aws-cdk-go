//go:build no_runtime_type_checking

package awscdk

// Building without runtime type checking enabled, so all the below just return nil

func validateBitrate_BpsParameters(amount *float64) error {
	return nil
}

func validateBitrate_GbpsParameters(amount *float64) error {
	return nil
}

func validateBitrate_KbpsParameters(amount *float64) error {
	return nil
}

func validateBitrate_MbpsParameters(amount *float64) error {
	return nil
}


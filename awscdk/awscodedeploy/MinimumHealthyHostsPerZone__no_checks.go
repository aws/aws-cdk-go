//go:build no_runtime_type_checking

package awscodedeploy

// Building without runtime type checking enabled, so all the below just return nil

func validateMinimumHealthyHostsPerZone_CountParameters(value *float64) error {
	return nil
}

func validateMinimumHealthyHostsPerZone_PercentageParameters(value *float64) error {
	return nil
}


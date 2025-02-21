//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func validateStats_PParameters(percentile *float64) error {
	return nil
}

func validateStats_PercentileParameters(percentile *float64) error {
	return nil
}

func validateStats_PercentileRankParameters(v1 *float64) error {
	return nil
}

func validateStats_PrParameters(v1 *float64) error {
	return nil
}

func validateStats_TcParameters(p1 *float64) error {
	return nil
}

func validateStats_TmParameters(p1 *float64) error {
	return nil
}

func validateStats_TrimmedCountParameters(p1 *float64) error {
	return nil
}

func validateStats_TrimmedMeanParameters(p1 *float64) error {
	return nil
}

func validateStats_TrimmedSumParameters(p1 *float64) error {
	return nil
}

func validateStats_TsParameters(p1 *float64) error {
	return nil
}

func validateStats_WinsorizedMeanParameters(p1 *float64) error {
	return nil
}

func validateStats_WmParameters(p1 *float64) error {
	return nil
}


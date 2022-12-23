//go:build !no_runtime_type_checking

package awscloudwatch

import (
	"fmt"
)

func validateStats_PParameters(percentile *float64) error {
	if percentile == nil {
		return fmt.Errorf("parameter percentile is required, but nil was provided")
	}

	return nil
}

func validateStats_PercentileParameters(percentile *float64) error {
	if percentile == nil {
		return fmt.Errorf("parameter percentile is required, but nil was provided")
	}

	return nil
}

func validateStats_PercentileRankParameters(v1 *float64) error {
	if v1 == nil {
		return fmt.Errorf("parameter v1 is required, but nil was provided")
	}

	return nil
}

func validateStats_PrParameters(v1 *float64) error {
	if v1 == nil {
		return fmt.Errorf("parameter v1 is required, but nil was provided")
	}

	return nil
}

func validateStats_TcParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_TmParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_TrimmedCountParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_TrimmedMeanParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_TrimmedSumParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_TsParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_WinsorizedMeanParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}

func validateStats_WmParameters(p1 *float64) error {
	if p1 == nil {
		return fmt.Errorf("parameter p1 is required, but nil was provided")
	}

	return nil
}


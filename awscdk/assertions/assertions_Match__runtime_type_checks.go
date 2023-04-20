//go:build !no_runtime_type_checking

package assertions

import (
	"fmt"
)

func validateMatch_ArrayEqualsParameters(pattern *[]interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_ArrayWithParameters(pattern *[]interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_ExactParameters(pattern interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_NotParameters(pattern interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_ObjectEqualsParameters(pattern *map[string]interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_ObjectLikeParameters(pattern *map[string]interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_SerializedJsonParameters(pattern interface{}) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}

func validateMatch_StringLikeRegexpParameters(pattern *string) error {
	if pattern == nil {
		return fmt.Errorf("parameter pattern is required, but nil was provided")
	}

	return nil
}


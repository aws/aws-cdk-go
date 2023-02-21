//go:build !no_runtime_type_checking

package awslambda

import (
	"fmt"
)

func validateFilterRule_BeginsWithParameters(elem *string) error {
	if elem == nil {
		return fmt.Errorf("parameter elem is required, but nil was provided")
	}

	return nil
}

func validateFilterRule_BetweenParameters(first *float64, second *float64) error {
	if first == nil {
		return fmt.Errorf("parameter first is required, but nil was provided")
	}

	if second == nil {
		return fmt.Errorf("parameter second is required, but nil was provided")
	}

	return nil
}

func validateFilterRule_IsEqualParameters(item interface{}) error {
	if item == nil {
		return fmt.Errorf("parameter item is required, but nil was provided")
	}
	switch item.(type) {
	case *string:
		// ok
	case string:
		// ok
	case *float64:
		// ok
	case float64:
		// ok
	case *int:
		// ok
	case int:
		// ok
	case *uint:
		// ok
	case uint:
		// ok
	case *int8:
		// ok
	case int8:
		// ok
	case *int16:
		// ok
	case int16:
		// ok
	case *int32:
		// ok
	case int32:
		// ok
	case *int64:
		// ok
	case int64:
		// ok
	case *uint8:
		// ok
	case uint8:
		// ok
	case *uint16:
		// ok
	case uint16:
		// ok
	case *uint32:
		// ok
	case uint32:
		// ok
	case *uint64:
		// ok
	case uint64:
		// ok
	default:
		return fmt.Errorf("parameter item must be one of the allowed types: *string, *float64; received %#v (a %T)", item, item)
	}

	return nil
}

func validateFilterRule_NotEqualsParameters(elem *string) error {
	if elem == nil {
		return fmt.Errorf("parameter elem is required, but nil was provided")
	}

	return nil
}


//go:build !no_runtime_type_checking

package awsbedrockagentcorealpha

import (
	"fmt"
)

func (c *jsiiProxy_ConditionalAttributeAccessor) validateContainsParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateEqualToParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
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
	case *bool:
		// ok
	case bool:
		// ok
	default:
		return fmt.Errorf("parameter value must be one of the allowed types: *string, *float64, *bool; received %#v (a %T)", value, value)
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateGreaterThanParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateGreaterThanOrEqualToParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateIsInParameters(values *[]interface{}) error {
	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}
	for idx_89445e, v := range *values {
		switch v.(type) {
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
			return fmt.Errorf("parameter values[%#v] must be one of the allowed types: *string, *float64; received %#v (a %T)", idx_89445e, v, v)
		}
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateIsInRangeParameters(ipRange *string) error {
	if ipRange == nil {
		return fmt.Errorf("parameter ipRange is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateLessThanParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateLessThanOrEqualToParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ConditionalAttributeAccessor) validateNotEqualToParameters(value interface{}) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}
	switch value.(type) {
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
	case *bool:
		// ok
	case bool:
		// ok
	default:
		return fmt.Errorf("parameter value must be one of the allowed types: *string, *float64, *bool; received %#v (a %T)", value, value)
	}

	return nil
}

func validateNewConditionalAttributeAccessorParameters(path *string, parent ConditionalPolicyStatement, conditionBuilder ConditionBuilder) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if parent == nil {
		return fmt.Errorf("parameter parent is required, but nil was provided")
	}

	if conditionBuilder == nil {
		return fmt.Errorf("parameter conditionBuilder is required, but nil was provided")
	}

	return nil
}


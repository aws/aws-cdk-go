//go:build !no_runtime_type_checking

package awsbedrockagentcore

import (
	"fmt"
)

func validatePolicyCondition_AllOfParameters(conditions *[]PolicyCondition) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_AnyOfParameters(conditions *[]PolicyCondition) error {
	if conditions == nil {
		return fmt.Errorf("parameter conditions is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_BooleanEqualsParameters(attribute PolicyAttribute, value *bool) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_IpInRangeParameters(attribute PolicyAttribute, cidr *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if cidr == nil {
		return fmt.Errorf("parameter cidr is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberEqualsParameters(attribute PolicyAttribute, value *float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberGreaterThanParameters(attribute PolicyAttribute, value *float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberGreaterThanOrEqualsParameters(attribute PolicyAttribute, value *float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberInParameters(attribute PolicyAttribute, values *[]*float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberLessThanParameters(attribute PolicyAttribute, value *float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberLessThanOrEqualsParameters(attribute PolicyAttribute, value *float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_NumberNotEqualsParameters(attribute PolicyAttribute, value *float64) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_SetContainsParameters(attribute PolicyAttribute, value *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_StringEqualsParameters(attribute PolicyAttribute, value *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_StringInParameters(attribute PolicyAttribute, values *[]*string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validatePolicyCondition_StringNotEqualsParameters(attribute PolicyAttribute, value *string) error {
	if attribute == nil {
		return fmt.Errorf("parameter attribute is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}


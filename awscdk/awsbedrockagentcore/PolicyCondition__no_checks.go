//go:build no_runtime_type_checking

package awsbedrockagentcore

// Building without runtime type checking enabled, so all the below just return nil

func validatePolicyCondition_AllOfParameters(conditions *[]PolicyCondition) error {
	return nil
}

func validatePolicyCondition_AnyOfParameters(conditions *[]PolicyCondition) error {
	return nil
}

func validatePolicyCondition_BooleanEqualsParameters(attribute PolicyAttribute, value *bool) error {
	return nil
}

func validatePolicyCondition_IpInRangeParameters(attribute PolicyAttribute, cidr *string) error {
	return nil
}

func validatePolicyCondition_NumberEqualsParameters(attribute PolicyAttribute, value *float64) error {
	return nil
}

func validatePolicyCondition_NumberGreaterThanParameters(attribute PolicyAttribute, value *float64) error {
	return nil
}

func validatePolicyCondition_NumberGreaterThanOrEqualsParameters(attribute PolicyAttribute, value *float64) error {
	return nil
}

func validatePolicyCondition_NumberInParameters(attribute PolicyAttribute, values *[]*float64) error {
	return nil
}

func validatePolicyCondition_NumberLessThanParameters(attribute PolicyAttribute, value *float64) error {
	return nil
}

func validatePolicyCondition_NumberLessThanOrEqualsParameters(attribute PolicyAttribute, value *float64) error {
	return nil
}

func validatePolicyCondition_NumberNotEqualsParameters(attribute PolicyAttribute, value *float64) error {
	return nil
}

func validatePolicyCondition_SetContainsParameters(attribute PolicyAttribute, value *string) error {
	return nil
}

func validatePolicyCondition_StringEqualsParameters(attribute PolicyAttribute, value *string) error {
	return nil
}

func validatePolicyCondition_StringInParameters(attribute PolicyAttribute, values *[]*string) error {
	return nil
}

func validatePolicyCondition_StringNotEqualsParameters(attribute PolicyAttribute, value *string) error {
	return nil
}


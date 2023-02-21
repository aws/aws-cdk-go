//go:build no_runtime_type_checking

package awsevents

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Match) validateResolveParameters(context awscdk.IResolveContext) error {
	return nil
}

func validateMatch_AnythingButPrefixParameters(prefix *string) error {
	return nil
}

func validateMatch_CidrParameters(range_ *string) error {
	return nil
}

func validateMatch_EqualParameters(value *float64) error {
	return nil
}

func validateMatch_EqualsIgnoreCaseParameters(value *string) error {
	return nil
}

func validateMatch_ExactStringParameters(value *string) error {
	return nil
}

func validateMatch_GreaterThanParameters(value *float64) error {
	return nil
}

func validateMatch_GreaterThanOrEqualParameters(value *float64) error {
	return nil
}

func validateMatch_IntervalParameters(lower *float64, upper *float64) error {
	return nil
}

func validateMatch_IpAddressRangeParameters(range_ *string) error {
	return nil
}

func validateMatch_LessThanParameters(value *float64) error {
	return nil
}

func validateMatch_LessThanOrEqualParameters(value *float64) error {
	return nil
}

func validateMatch_PrefixParameters(value *string) error {
	return nil
}

func validateMatch_SuffixParameters(value *string) error {
	return nil
}


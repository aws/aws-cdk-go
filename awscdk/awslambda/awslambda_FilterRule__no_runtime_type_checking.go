//go:build no_runtime_type_checking

package awslambda

// Building without runtime type checking enabled, so all the below just return nil

func validateFilterRule_BeginsWithParameters(elem *string) error {
	return nil
}

func validateFilterRule_BetweenParameters(first *float64, second *float64) error {
	return nil
}

func validateFilterRule_IsEqualParameters(item interface{}) error {
	return nil
}

func validateFilterRule_NotEqualsParameters(elem *string) error {
	return nil
}


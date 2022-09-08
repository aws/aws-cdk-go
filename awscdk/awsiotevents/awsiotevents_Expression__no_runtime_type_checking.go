//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func validateExpression_AndParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_CurrentInputParameters(input IInput) error {
	return nil
}

func validateExpression_EqParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_FromStringParameters(value *string) error {
	return nil
}

func validateExpression_GtParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_GteParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_InputAttributeParameters(input IInput, path *string) error {
	return nil
}

func validateExpression_LtParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_LteParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_NeqParameters(left Expression, right Expression) error {
	return nil
}

func validateExpression_OrParameters(left Expression, right Expression) error {
	return nil
}


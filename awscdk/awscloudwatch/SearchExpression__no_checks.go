//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SearchExpression) validateWithParameters(props *SearchExpressionOptions) error {
	return nil
}

func validateNewSearchExpressionParameters(props *SearchExpressionProps) error {
	return nil
}


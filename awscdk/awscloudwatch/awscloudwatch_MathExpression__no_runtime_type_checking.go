//go:build no_runtime_type_checking

package awscloudwatch

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_MathExpression) validateCreateAlarmParameters(scope constructs.Construct, id *string, props *CreateAlarmOptions) error {
	return nil
}

func (m *jsiiProxy_MathExpression) validateWithParameters(props *MathExpressionOptions) error {
	return nil
}

func validateNewMathExpressionParameters(props *MathExpressionProps) error {
	return nil
}


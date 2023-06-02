//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (l *jsiiProxy_Lambda) validateBindParameters(rule awsses.IReceiptRule) error {
	return nil
}

func validateNewLambdaParameters(props *LambdaProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkMail) validateBindParameters(receiptRule awsses.IReceiptRule) error {
	return nil
}

func validateNewWorkMailParameters(props *WorkMailProps) error {
	return nil
}


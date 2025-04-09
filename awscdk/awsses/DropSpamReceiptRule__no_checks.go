//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func validateDropSpamReceiptRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDropSpamReceiptRuleParameters(scope constructs.Construct, id *string, props *DropSpamReceiptRuleProps) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DropSpamReceiptRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DropSpamReceiptRule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDropSpamReceiptRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDropSpamReceiptRuleParameters(scope constructs.Construct, id *string, props *DropSpamReceiptRuleProps) error {
	return nil
}


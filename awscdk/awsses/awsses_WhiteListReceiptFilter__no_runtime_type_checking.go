//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WhiteListReceiptFilter) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (w *jsiiProxy_WhiteListReceiptFilter) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateWhiteListReceiptFilter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewWhiteListReceiptFilterParameters(scope constructs.Construct, id *string, props *WhiteListReceiptFilterProps) error {
	return nil
}


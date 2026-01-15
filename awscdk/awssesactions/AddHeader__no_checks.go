//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AddHeader) validateBindParameters(receiptRule interfacesawsses.IReceiptRuleRef) error {
	return nil
}

func validateNewAddHeaderParameters(props *AddHeaderProps) error {
	return nil
}


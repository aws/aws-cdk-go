//go:build no_runtime_type_checking

package awssesactions

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_Bounce) validateBindParameters(receiptRule interfacesawsses.IReceiptRuleRef) error {
	return nil
}

func validateNewBounceParameters(props *BounceProps) error {
	return nil
}


//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReceiptFilter) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ReceiptFilter) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ReceiptFilter) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_ReceiptFilter) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_ReceiptFilter) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateReceiptFilter_IsConstructParameters(x interface{}) error {
	return nil
}

func validateReceiptFilter_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewReceiptFilterParameters(scope constructs.Construct, id *string, props *ReceiptFilterProps) error {
	return nil
}


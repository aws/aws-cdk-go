//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ReceiptRule) validateAddActionParameters(action IReceiptRuleAction) error {
	return nil
}

func (r *jsiiProxy_ReceiptRule) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ReceiptRule) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ReceiptRule) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_ReceiptRule) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_ReceiptRule) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateReceiptRule_FromReceiptRuleNameParameters(scope constructs.Construct, id *string, receiptRuleName *string) error {
	return nil
}

func validateReceiptRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateReceiptRule_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewReceiptRuleParameters(scope constructs.Construct, id *string, props *ReceiptRuleProps) error {
	return nil
}


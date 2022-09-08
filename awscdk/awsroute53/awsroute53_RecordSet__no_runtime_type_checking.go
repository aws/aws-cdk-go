//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsroute53

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RecordSet) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_RecordSet) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_RecordSet) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (r *jsiiProxy_RecordSet) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (r *jsiiProxy_RecordSet) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateRecordSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRecordSet_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewRecordSetParameters(scope constructs.Construct, id *string, props *RecordSetProps) error {
	return nil
}


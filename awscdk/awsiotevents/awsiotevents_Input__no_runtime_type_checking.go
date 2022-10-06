//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsiotevents

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_Input) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_Input) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (i *jsiiProxy_Input) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (i *jsiiProxy_Input) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_Input) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_Input) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (i *jsiiProxy_Input) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateInput_FromInputNameParameters(scope constructs.Construct, id *string, inputName *string) error {
	return nil
}

func validateInput_IsConstructParameters(x interface{}) error {
	return nil
}

func validateInput_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewInputParameters(scope constructs.Construct, id *string, props *InputProps) error {
	return nil
}


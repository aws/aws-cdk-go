//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (o *jsiiProxy_OptionGroup) validateAddConfigurationParameters(configuration *OptionConfiguration) error {
	return nil
}

func (o *jsiiProxy_OptionGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (o *jsiiProxy_OptionGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (o *jsiiProxy_OptionGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (o *jsiiProxy_OptionGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (o *jsiiProxy_OptionGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateOptionGroup_FromOptionGroupNameParameters(scope constructs.Construct, id *string, optionGroupName *string) error {
	return nil
}

func validateOptionGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateOptionGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewOptionGroupParameters(scope constructs.Construct, id *string, props *OptionGroupProps) error {
	return nil
}


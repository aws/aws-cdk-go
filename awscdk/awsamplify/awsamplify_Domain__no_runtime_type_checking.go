//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsamplify

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Domain) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Domain) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Domain) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_Domain) validateMapRootParameters(branch IBranch) error {
	return nil
}

func (d *jsiiProxy_Domain) validateMapSubDomainParameters(branch IBranch) error {
	return nil
}

func (d *jsiiProxy_Domain) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_Domain) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDomain_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDomain_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDomainParameters(scope constructs.Construct, id *string, props *DomainProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_Addon) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_Addon) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAddon_FromAddonArnParameters(scope constructs.Construct, id *string, addonArn *string) error {
	return nil
}

func validateAddon_FromAddonAttributesParameters(scope constructs.Construct, id *string, attrs *AddonAttributes) error {
	return nil
}

func validateAddon_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAddon_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAddon_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAddonParameters(scope constructs.Construct, id *string, props *AddonProps) error {
	return nil
}


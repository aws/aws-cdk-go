//go:build no_runtime_type_checking

package awsec2

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PrefixList) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (p *jsiiProxy_PrefixList) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (p *jsiiProxy_PrefixList) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validatePrefixList_FromPrefixListIdParameters(scope constructs.Construct, id *string, prefixListId *string) error {
	return nil
}

func validatePrefixList_IsConstructParameters(x interface{}) error {
	return nil
}

func validatePrefixList_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validatePrefixList_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewPrefixListParameters(scope constructs.Construct, id *string, props *PrefixListProps) error {
	return nil
}


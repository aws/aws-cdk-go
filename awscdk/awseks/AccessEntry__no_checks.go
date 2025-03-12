//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AccessEntry) validateAddAccessPoliciesParameters(newAccessPolicies *[]IAccessPolicy) error {
	return nil
}

func (a *jsiiProxy_AccessEntry) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (a *jsiiProxy_AccessEntry) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (a *jsiiProxy_AccessEntry) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateAccessEntry_FromAccessEntryAttributesParameters(scope constructs.Construct, id *string, attrs *AccessEntryAttributes) error {
	return nil
}

func validateAccessEntry_IsConstructParameters(x interface{}) error {
	return nil
}

func validateAccessEntry_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateAccessEntry_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewAccessEntryParameters(scope constructs.Construct, id *string, props *AccessEntryProps) error {
	return nil
}


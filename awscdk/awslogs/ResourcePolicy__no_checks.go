//go:build no_runtime_type_checking

package awslogs

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_ResourcePolicy) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (r *jsiiProxy_ResourcePolicy) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateResourcePolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateResourcePolicy_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateResourcePolicy_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewResourcePolicyParameters(scope constructs.Construct, id *string, props *ResourcePolicyProps) error {
	return nil
}


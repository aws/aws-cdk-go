//go:build no_runtime_type_checking

package awscdkmskalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClusterBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateClusterBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateClusterBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClusterBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}


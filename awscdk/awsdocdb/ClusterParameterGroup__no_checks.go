//go:build no_runtime_type_checking

package awsdocdb

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterParameterGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClusterParameterGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClusterParameterGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateClusterParameterGroup_FromParameterGroupNameParameters(scope constructs.Construct, id *string, parameterGroupName *string) error {
	return nil
}

func validateClusterParameterGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterParameterGroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateClusterParameterGroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClusterParameterGroupParameters(scope constructs.Construct, id *string, props *ClusterParameterGroupProps) error {
	return nil
}


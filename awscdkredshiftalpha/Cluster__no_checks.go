//go:build no_runtime_type_checking

package awscdkredshiftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Cluster) validateAddDefaultIamRoleParameters(defaultIamRole awsiam.IRole) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddIamRoleParameters(role awsiam.IRole) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateAddToParameterGroupParameters(name *string, value *string) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_Cluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateCluster_FromClusterAttributesParameters(scope constructs.Construct, id *string, attrs *ClusterAttributes) error {
	return nil
}

func validateCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateCluster_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateCluster_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewClusterParameters(scope constructs.Construct, id *string, props *ClusterProps) error {
	return nil
}


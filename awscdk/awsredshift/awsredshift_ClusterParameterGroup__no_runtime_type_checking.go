//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsredshift

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

func (c *jsiiProxy_ClusterParameterGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_ClusterParameterGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateClusterParameterGroup_FromClusterParameterGroupNameParameters(scope constructs.Construct, id *string, clusterParameterGroupName *string) error {
	return nil
}

func validateClusterParameterGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterParameterGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewClusterParameterGroupParameters(scope constructs.Construct, id *string, props *ClusterParameterGroupProps) error {
	return nil
}


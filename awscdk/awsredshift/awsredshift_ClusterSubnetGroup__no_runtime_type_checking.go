//go:build no_runtime_type_checking

package awsredshift

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterSubnetGroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (c *jsiiProxy_ClusterSubnetGroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (c *jsiiProxy_ClusterSubnetGroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (c *jsiiProxy_ClusterSubnetGroup) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (c *jsiiProxy_ClusterSubnetGroup) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateClusterSubnetGroup_FromClusterSubnetGroupNameParameters(scope constructs.Construct, id *string, clusterSubnetGroupName *string) error {
	return nil
}

func validateClusterSubnetGroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateClusterSubnetGroup_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewClusterSubnetGroupParameters(scope constructs.Construct, id *string, props *ClusterSubnetGroupProps) error {
	return nil
}


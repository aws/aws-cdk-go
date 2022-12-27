//go:build no_runtime_type_checking

package awseks

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_Nodegroup) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (n *jsiiProxy_Nodegroup) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (n *jsiiProxy_Nodegroup) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func validateNodegroup_FromNodegroupNameParameters(scope constructs.Construct, id *string, nodegroupName *string) error {
	return nil
}

func validateNodegroup_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNodegroup_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNodegroup_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewNodegroupParameters(scope constructs.Construct, id *string, props *NodegroupProps) error {
	return nil
}


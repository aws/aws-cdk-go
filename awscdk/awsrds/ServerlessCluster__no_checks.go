//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerlessCluster) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateAddRotationSingleUserParameters(options *RotationSingleUserOptions) error {
	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_ServerlessCluster) validateGrantDataApiAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateServerlessCluster_FromServerlessClusterAttributesParameters(scope constructs.Construct, id *string, attrs *ServerlessClusterAttributes) error {
	return nil
}

func validateServerlessCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServerlessCluster_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateServerlessCluster_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServerlessClusterParameters(scope constructs.Construct, id *string, props *ServerlessClusterProps) error {
	return nil
}


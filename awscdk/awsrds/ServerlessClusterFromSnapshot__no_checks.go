//go:build no_runtime_type_checking

package awsrds

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServerlessClusterFromSnapshot) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (s *jsiiProxy_ServerlessClusterFromSnapshot) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (s *jsiiProxy_ServerlessClusterFromSnapshot) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (s *jsiiProxy_ServerlessClusterFromSnapshot) validateGrantDataApiAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateServerlessClusterFromSnapshot_IsConstructParameters(x interface{}) error {
	return nil
}

func validateServerlessClusterFromSnapshot_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateServerlessClusterFromSnapshot_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewServerlessClusterFromSnapshotParameters(scope constructs.Construct, id *string, props *ServerlessClusterFromSnapshotProps) error {
	return nil
}


//go:build no_runtime_type_checking

package awsdocdb

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseCluster) validateAddRotationMultiUserParameters(id *string, options *RotationMultiUserOptions) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DatabaseCluster) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDatabaseCluster_FromDatabaseClusterAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) error {
	return nil
}

func validateDatabaseCluster_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseCluster_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDatabaseClusterParameters(scope constructs.Construct, id *string, props *DatabaseClusterProps) error {
	return nil
}


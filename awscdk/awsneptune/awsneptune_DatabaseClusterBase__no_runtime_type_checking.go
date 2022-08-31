//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsneptune

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseClusterBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateGrantConnectParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DatabaseClusterBase) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDatabaseClusterBase_FromDatabaseClusterAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseClusterAttributes) error {
	return nil
}

func validateDatabaseClusterBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseClusterBase_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDatabaseClusterBaseParameters(scope constructs.Construct, id *string, props *awscdk.ResourceProps) error {
	return nil
}


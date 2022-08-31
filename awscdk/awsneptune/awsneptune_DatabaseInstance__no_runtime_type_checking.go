//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsneptune

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DatabaseInstance) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_DatabaseInstance) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDatabaseInstance_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstance_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceParameters(scope constructs.Construct, id *string, props *DatabaseInstanceProps) error {
	return nil
}


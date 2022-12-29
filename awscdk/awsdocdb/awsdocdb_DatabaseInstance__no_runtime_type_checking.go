//go:build no_runtime_type_checking

package awsdocdb

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

func validateDatabaseInstance_FromDatabaseInstanceAttributesParameters(scope constructs.Construct, id *string, attrs *DatabaseInstanceAttributes) error {
	return nil
}

func validateDatabaseInstance_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabaseInstance_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateDatabaseInstance_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewDatabaseInstanceParameters(scope constructs.Construct, id *string, props *DatabaseInstanceProps) error {
	return nil
}


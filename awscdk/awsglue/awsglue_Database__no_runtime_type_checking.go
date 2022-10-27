//go:build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Database) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (d *jsiiProxy_Database) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (d *jsiiProxy_Database) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (d *jsiiProxy_Database) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateDatabase_FromDatabaseArnParameters(scope constructs.Construct, id *string, databaseArn *string) error {
	return nil
}

func validateDatabase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateDatabase_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewDatabaseParameters(scope constructs.Construct, id *string, props *DatabaseProps) error {
	return nil
}


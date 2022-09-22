//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsglue

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Table) validateAddPartitionIndexParameters(index *PartitionIndex) error {
	return nil
}

func (t *jsiiProxy_Table) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Table) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_Table) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantToUnderlyingResourcesParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_Table) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (t *jsiiProxy_Table) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateTable_FromTableArnParameters(scope constructs.Construct, id *string, tableArn *string) error {
	return nil
}

func validateTable_FromTableAttributesParameters(scope constructs.Construct, id *string, attrs *TableAttributes) error {
	return nil
}

func validateTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTable_IsResourceParameters(construct awscdk.IConstruct) error {
	return nil
}

func validateNewTableParameters(scope constructs.Construct, id *string, props *TableProps) error {
	return nil
}


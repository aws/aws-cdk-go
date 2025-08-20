//go:build no_runtime_type_checking

package awscdkgluealpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableBase) validateAddPartitionIndexParameters(index *PartitionIndex) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGetResourceArnAttributeParameters(arnAttr *string, arnComponents *awscdk.ArnComponents) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGetResourceNameAttributeParameters(nameAttr *string) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantReadWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantToUnderlyingResourcesParameters(grantee awsiam.IGrantable, actions *[]*string) error {
	return nil
}

func (t *jsiiProxy_TableBase) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateTableBase_FromTableArnParameters(scope constructs.Construct, id *string, tableArn *string) error {
	return nil
}

func validateTableBase_FromTableAttributesParameters(scope constructs.Construct, id *string, attrs *TableAttributes) error {
	return nil
}

func validateTableBase_IsConstructParameters(x interface{}) error {
	return nil
}

func validateTableBase_IsOwnedResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateTableBase_IsResourceParameters(construct constructs.IConstruct) error {
	return nil
}

func validateNewTableBaseParameters(scope constructs.Construct, id *string, props *TableBaseProps) error {
	return nil
}


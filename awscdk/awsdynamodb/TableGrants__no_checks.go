//go:build no_runtime_type_checking

package awsdynamodb

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TableGrants) validateActionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableGrants) validateFullAccessParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableGrants) validateReadDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableGrants) validateReadWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (t *jsiiProxy_TableGrants) validateWriteDataParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateNewTableGrantsParameters(props *TableGrantsProps) error {
	return nil
}


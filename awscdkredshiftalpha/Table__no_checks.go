//go:build no_runtime_type_checking

package awscdkredshiftalpha

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_Table) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (t *jsiiProxy_Table) validateGrantParameters(user IUser) error {
	return nil
}

func validateTable_FromTableAttributesParameters(scope constructs.Construct, id *string, attrs *TableAttributes) error {
	return nil
}

func validateTable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTableParameters(scope constructs.Construct, id *string, props *TableProps) error {
	return nil
}


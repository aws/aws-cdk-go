//go:build no_runtime_type_checking

package awsredshift

// Building without runtime type checking enabled, so all the below just return nil

func (u *jsiiProxy_User) validateAddTablePrivilegesParameters(table ITable) error {
	return nil
}

func (u *jsiiProxy_User) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (u *jsiiProxy_User) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (u *jsiiProxy_User) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateUser_FromUserAttributesParameters(scope constructs.Construct, id *string, attrs *UserAttributes) error {
	return nil
}

func validateUser_IsConstructParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_User) validateSetDatabasePropsParameters(val *DatabaseOptions) error {
	return nil
}

func validateNewUserParameters(scope constructs.Construct, id *string, props *UserProps) error {
	return nil
}


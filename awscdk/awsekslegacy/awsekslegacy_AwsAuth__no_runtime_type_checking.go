//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package awsekslegacy

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsAuth) validateAddAccountParameters(accountId *string) error {
	return nil
}

func (a *jsiiProxy_AwsAuth) validateAddMastersRoleParameters(role awsiam.IRole) error {
	return nil
}

func (a *jsiiProxy_AwsAuth) validateAddRoleMappingParameters(role awsiam.IRole, mapping *Mapping) error {
	return nil
}

func (a *jsiiProxy_AwsAuth) validateAddUserMappingParameters(user awsiam.IUser, mapping *Mapping) error {
	return nil
}

func (a *jsiiProxy_AwsAuth) validateOnSynthesizeParameters(session constructs.ISynthesisSession) error {
	return nil
}

func (a *jsiiProxy_AwsAuth) validateSynthesizeParameters(session awscdk.ISynthesisSession) error {
	return nil
}

func validateAwsAuth_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsAuthParameters(scope awscdk.Construct, id *string, props *AwsAuthProps) error {
	return nil
}


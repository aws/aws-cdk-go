//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IAlias) validateAddAliasParameters(alias *string) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateApplyRemovalPolicyParameters(policy awscdk.RemovalPolicy) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantEncryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantEncryptDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantGenerateMacParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantSignParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantSignVerifyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantVerifyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IAlias) validateGrantVerifyMacParameters(grantee awsiam.IGrantable) error {
	return nil
}


//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IKey) validateAddAliasParameters(alias *string) error {
	return nil
}

func (i *jsiiProxy_IKey) validateAddToResourcePolicyParameters(statement awsiam.PolicyStatement) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantEncryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantEncryptDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantGenerateMacParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantSignParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantSignVerifyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantVerifyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IKey) validateGrantVerifyMacParameters(grantee awsiam.IGrantable) error {
	return nil
}


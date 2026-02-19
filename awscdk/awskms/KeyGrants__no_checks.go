//go:build no_runtime_type_checking

package awskms

// Building without runtime type checking enabled, so all the below just return nil

func (k *jsiiProxy_KeyGrants) validateActionsParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateAdminParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateEncryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateEncryptDecryptParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateGenerateMacParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateSignParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateSignVerifyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateVerifyParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (k *jsiiProxy_KeyGrants) validateVerifyMacParameters(grantee awsiam.IGrantable) error {
	return nil
}

func validateKeyGrants_FromKeyParameters(resource interfacesawskms.IKeyRef) error {
	return nil
}


//go:build no_runtime_type_checking

package awsses

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IEmailIdentity) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IEmailIdentity) validateGrantSendEmailParameters(grantee awsiam.IGrantable) error {
	return nil
}


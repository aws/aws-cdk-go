//go:build no_runtime_type_checking

package awsssm

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IParameter) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IParameter) validateGrantWriteParameters(grantee awsiam.IGrantable) error {
	return nil
}


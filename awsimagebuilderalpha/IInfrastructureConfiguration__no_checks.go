//go:build no_runtime_type_checking

package awsimagebuilderalpha

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IInfrastructureConfiguration) validateGrantParameters(grantee awsiam.IGrantable) error {
	return nil
}

func (i *jsiiProxy_IInfrastructureConfiguration) validateGrantReadParameters(grantee awsiam.IGrantable) error {
	return nil
}

